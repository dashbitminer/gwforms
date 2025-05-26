var signaturePad;
  
  function signiture(){
    var input = document.getElementById("Signature1draw");
    signaturePad = new SignaturePad(input);

  }

  function clearSigniture(){
    signaturePad.clear();
  }

  document.addEventListener('DOMContentLoaded', function() {
    signiture();

      var fecha = new Date();
      var dia = ('0' + fecha.getDate()).slice(-2); // Asegura el formato de dos dígitos
      var mes = ('0' + (fecha.getMonth() + 1)).slice(-2); // Asegura el formato de dos dígitos y corrige el índice del mes
      var ano = fecha.getFullYear();
      var fechaFormateada = dia+ '-' + mes + '-' + ano ;
      document.getElementById('fechaActual').textContent = fechaFormateada;
      var fechaFormateada = ano + '-' + mes + '-' + dia; // Formato ajustado para <input type="date">
      document.getElementById('fechafirma').value = fechaFormateada;
  });


function guardar(event) {
    event.preventDefault();
    var form = document.getElementById("form");
    var invalidElements = form.querySelectorAll('.is-invalid');
    if (!form.checkValidity() || invalidElements.length > 0) {
        Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Por favor, completa el formulario correctamente antes de enviar.',
        });
        return;
    }
    // Verifica si signaturePad está vacío
    if (typeof signaturePad !== "undefined" && signaturePad && signaturePad.isEmpty()) {
        Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Por favor, asegúrate de firmar antes de enviar.',
        });
        return; // Detiene la ejecución si no hay firma
    }

    var formData = new FormData(form);
    // Añade la firma al formData si signaturePad está definido y no está vacío
    if (typeof signaturePad !== "undefined" && signaturePad && !signaturePad.isEmpty()) {
        var firma1 = signaturePad.toDataURL();
        formData.append("Signature1draw", firma1);
    }


    // Espera a que descargarPDF se complete antes de proceder
    descargarPDF().then((pdfBlob) => {
        formData.append("pdfConsentimiento", pdfBlob, "Consentimiento.pdf");
        // Muestra una alerta de "Subiendo..."
        Swal.fire({
            title: 'Subiendo...',
            html: 'Por favor, espera un momento.',
            allowOutsideClick: false,
            onBeforeOpen: () => {
                Swal.showLoading()
            },
        });

        var xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function() {
            if (this.readyState == 4) {
                Swal.close(); // Cierra la alerta de "Subiendo..." cuando la petición finaliza
                if (this.status == 200) {
                    var respuesta = this.responseText;
                    if (respuesta == "1") {
                        Swal.fire({
                            title: 'Consentimiento registrado correctamente',
                            icon: 'success',
                            html: `
                                <p>El consentimiento se ha registrado correctamente.</p>
                                <input type="email" id="emailInput" class="swal2-input" placeholder="Ingrese su correo electrónico">
                            `,
                            showCancelButton: true,
                            confirmButtonText: 'Descargar PDF',
                            cancelButtonText: 'Cerrar',
                            showDenyButton: true,
                            denyButtonText: 'Enviar correo electrónico',
                        }).then((result) => {
                            if (result.isConfirmed) {
                                // Llamada a la función para descargar el PDF
                                descargarPDFlocal();
                            } else if (result.isDenied) {
                                // Obtener el correo electrónico ingresado
                                var email = document.getElementById('emailInput').value;
                                if (email) {
                                    // Llamada a la función para enviar el correo electrónico
                                    enviarCorreoElectronico(email);
                                } else {
                                    Swal.fire({
                                        icon: 'error',
                                        title: 'Oops...',
                                        text: 'Por favor, ingrese un correo electrónico válido.',
                                    });
                                }
                            } else {
                                // Redirige a la página de preinscripción después de que el usuario cierre el SweetAlert
                                window.location.href = "/salud/consentimiento/mx";
                            }
                        });
                    } else {
                        Swal.fire({
                            icon: 'error',
                            title: 'Oops...',
                            text: 'Error al enviar el consentimiento',
                        });
                    }
                }
            }
        };
        xhr.open("POST", "/salud/consentimiento/gt/post", true);
        xhr.send(formData);
    });
}
var formElements = form.querySelectorAll('input, textarea, select');
for (var i = 0; i < formElements.length; i++) {
    formElements[i].addEventListener('change', function() {
      if (this.validity.valid) {
        this.classList.remove('is-invalid');
        this.classList.add('is-valid');
      } else {
        this.classList.remove('is-valid');
        this.classList.add('is-invalid');
      }
    });
}

function enviarCorreoElectronico(email) {
  var formData = new FormData(form);

  // Añade la firma al formData si signaturePad está definido y no está vacío
  if (typeof signaturePad !== "undefined" && signaturePad && !signaturePad.isEmpty()) {
    var firma1 = signaturePad.toDataURL();
    formData.append("Signature1draw", firma1);
  }

  descargarPDF().then((pdfBlob) => {
    formData.append("pdfConsentimiento", pdfBlob, "Consentimiento.pdf");
    // Muestra una alerta de "Subiendo..."
    Swal.fire({
        title: 'Subiendo...',
        html: 'Por favor, espera un momento.',
        allowOutsideClick: false,
        onBeforeOpen: () => {
            Swal.showLoading()
        },
    });
  // Añade el correo electrónico al formData
  formData.append("email", email);

  // Realiza la solicitud AJAX para enviar el correo electrónico
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/salud/consentimiento/gt/correo", true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      Swal.fire({
        icon: 'success',
        title: 'Correo enviado',
        text: 'El correo electrónico se ha enviado correctamente.',
      });
    } else if (xhr.readyState == 4) {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: 'Hubo un problema al enviar el correo electrónico.',
      });
    }
  };
  xhr.send(formData);
});
}

function descargarPDF() {
  return new Promise((resolve, reject) => {
    // Ocultar elementos por ID al inicio
    document.getElementById('boton').style.display = 'none';
    document.getElementById('borrar1').style.display = 'none';


    window.jsPDF = window.jspdf.jsPDF;
    const doc = new jsPDF({
      orientation: "p",
      unit: "mm",
      format: [215.9, 279.4] // Tamaño de página carta
    });

    const elementsToCapture = [
      "#contentToPrint"
    ];

    const captureAndAddToPDF = (index = 0) => {
      if (index >= elementsToCapture.length) {
        // Cambio realizado aquí: Eliminar el uso de .then() con doc.output()
        const blob = doc.output('blob');
        // Mostrar nuevamente los elementos después de generar el PDF
        document.getElementById('boton').style.display = '';
        document.getElementById('borrar1').style.display = '';
        resolve(blob);
        return;
      }

      const elementSelector = elementsToCapture[index];
      const elementHTML = document.querySelector(elementSelector);

      html2canvas(elementHTML, { scale: 1 }).then(canvas => {
        const imgData = canvas.toDataURL('image/jpeg');
        const imgProps = doc.getImageProperties(imgData);
        const pdfWidth = doc.internal.pageSize.getWidth();
        const pdfHeight = doc.internal.pageSize.getHeight();
        const imgWidth = imgProps.width;
        const imgHeight = imgProps.height;
        let dimensions = calculateDimensions(imgWidth, imgHeight, pdfWidth, pdfHeight);

        if (index > 0) {
          doc.addPage();
        }

        doc.addImage(imgData, 'JPEG', dimensions.x, dimensions.y, dimensions.width, dimensions.height);
        captureAndAddToPDF(index + 1);
      });
    };

    captureAndAddToPDF();
  });
}

function descargarPDFlocal() {

    // Ocultar elementos por ID al inicio
    document.getElementById('boton').style.display = 'none';
    document.getElementById('borrar1').style.display = 'none';

  
    window.jsPDF = window.jspdf.jsPDF;
    const doc = new jsPDF({
      orientation: "p",
      unit: "mm",
      format: [215.9, 279.4] // Tamaño de página carta
    });
  
    const elementsToCapture = [
      "#contentToPrint"
    ];
  
    const captureAndAddToPDF = (index = 0) => {
      if (index >= elementsToCapture.length) {
        doc.save('Consentimiento.pdf');
        // Mostrar nuevamente los elementos después de guardar el PDF
        document.getElementById('boton').style.display = '';
        document.getElementById('borrar1').style.display = '';
        return;
      }
  
      const elementSelector = elementsToCapture[index];
      const elementHTML = document.querySelector(elementSelector);
  
      html2canvas(elementHTML, { scale: 1 }).then(canvas => {
        const imgData = canvas.toDataURL('image/jpeg');
        const imgProps = doc.getImageProperties(imgData);
        const pdfWidth = doc.internal.pageSize.getWidth();
        const pdfHeight = doc.internal.pageSize.getHeight();
        const imgWidth = imgProps.width;
        const imgHeight = imgProps.height;
        let dimensions = calculateDimensions(imgWidth, imgHeight, pdfWidth, pdfHeight);
  
        if (index > 0) {
          doc.addPage();
        }
  
        doc.addImage(imgData, 'JPEG', dimensions.x, dimensions.y, dimensions.width, dimensions.height);
        captureAndAddToPDF(index + 1);
      });
    };
  
    captureAndAddToPDF();
  
}

function calculateDimensions(imgWidth, imgHeight, pdfWidth, pdfHeight) {
  let imgAspectRatio = imgWidth / imgHeight;
  let pdfAspectRatio = pdfWidth / pdfHeight;
  let width, height, x, y;

  if (imgAspectRatio > pdfAspectRatio) {
    // La imagen es más ancha que el PDF, ajustar por ancho
    width = pdfWidth;
    height = pdfWidth / imgAspectRatio;
  } else {
    // La imagen es más alta que el PDF, ajustar por altura
    height = pdfHeight;
    width = pdfHeight * imgAspectRatio;
  }

  x = (pdfWidth - width) / 2;
  y = (pdfHeight - height) / 2;

  return { width, height, x, y };
}

function resizeCanvas(canvasId) {
  const canvas = document.getElementById(canvasId);
  const container = canvas.parentElement;
  canvas.width = container.offsetWidth;
  canvas.height = container.offsetHeight;
}

function clearSignature(canvasId) {
  const canvas = document.getElementById(canvasId);
  const context = canvas.getContext('2d');
  context.clearRect(0, 0, canvas.width, canvas.height);
}

window.addEventListener('resize', () => {
  resizeCanvas('Signature1draw');

});

// Inicializar el tamaño del canvas al cargar la página
document.addEventListener('DOMContentLoaded', () => {
  resizeCanvas('Signature1draw');

});