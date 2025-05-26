// Ejemplo de JavaScript inicial para deshabilitar el envío de formularios si hay campos no válidos
(function () {
    'use strict'
  
    // Obtener todos los formularios a los que queremos aplicar estilos de validación de Bootstrap personalizados
    var forms = document.querySelectorAll('.needs-validation')
  
    // Bucle sobre ellos y evitar el envío
    Array.prototype.slice.call(forms)
      .forEach(function (form) {
        form.addEventListener('submit', function (event) {
          if (!form.checkValidity()) {
            event.preventDefault()
            event.stopPropagation()
          }
  
          form.classList.add('was-validated')
        }, false)
      })
  })()


function persona(value){
var input = document.getElementById("nombres");
var input2 = document.getElementById("apellidos");
var input3 = document.getElementById("identidad");
var input4 = document.getElementById("identidadF");
var input5 = document.getElementById("identidadA");
var input6 = document.getElementById("NRC");
var input7 = document.getElementById("NRCF");
var input8 = document.getElementById("NRCA");
var input9 = document.getElementById("nombreComercial");
var input10 = document.getElementById("actividadEconomica");
var input11 = document.getElementById("nombreRepre");
var input12 = document.getElementById("razonSocial");
  
if (value ==='2'){
  document.getElementById('divnatural').classList.add('d-none');
  document.getElementById('divNRC').classList.remove('d-none');
  document.getElementById('divjuridica').classList.remove('d-none');
  input.required = false;
  input2.required = false;
  input3.required = false;
  input4.required = false;
  input5.required = false;
  input6.required = true;
  input7.required = true;
  input8.required = true;
  document.getElementById('divinscrita').classList.remove('d-none');
  input9.required = true;
  input10.required = true;
  input11.required = true;
  input12.required = true;
  document.getElementById('direccionlabel').innerHTML = "Dirección (Según registro de IVA):<span style='color: #FF7D81;'>**</span>";

  
} else if (value==='3'){
  document.getElementById('divnatural').classList.remove('d-none');
  document.getElementById('divNRC').classList.remove('d-none');
  document.getElementById('divjuridica').classList.add('d-none');
  input.required = true;
  input2.required = true;
  input3.required = true;
  input4.required = true;
  input5.required = true;
  input6.required = true;
  input7.required = true;
  input8.required = true;
  document.getElementById('divinscrita').classList.remove('d-none');
  input9.required = false;
  input10.required = false;
  input11.required = false;
  input12.required = false;
  document.getElementById('direccionlabel').innerHTML = "Dirección:<span style='color: #FF7D81;'>**</span>";

} else {
  document.getElementById('divnatural').classList.remove('d-none');
  document.getElementById('divjuridica').classList.add('d-none');
  document.getElementById('divNRC').classList.add('d-none');
  
  input.required = true;
  input2.required = true;
  input3.required = true;
  input4.required = true;
  input5.required = true;
  input6.required = false;
  input7.required = false;
  input8.required = false;

  document.getElementById('divinscrita').classList.add('d-none');
  input9.required = false;
  input10.required = false;
  input11.required = false;
  input12.required = false;
  document.getElementById('direccionlabel').innerHTML = "Dirección:<span style='color: #FF7D81;'>**</span>";
  }
}

function contribuyente() {
  var select = document.getElementById("tipoPersona");
  var hiddenDiv = document.getElementById("ocultarContribuyente");
  var input4 = document.getElementById("tipoContribuyente");

  if (select.value === "1") {
    hiddenDiv.classList.add("d-none");
    input4.required = false;
  } else {
    hiddenDiv.classList.remove("d-none");
    input4.required = true;
    input4.value = ""
  }
}

/* function inscrita(value){
  var input6 = document.getElementById("nombreComercial");
  var input7 = document.getElementById("actividadEconomica");
  
  if (value =='2'||value =='3'){
    document.getElementById('divinscrita').classList.remove('d-none');
    input6.required = true;
    input7.required = true;
  }
  else{
    document.getElementById('divinscrita').classList.add('d-none');
    input6.required = false;
    input7.required = false;
  }
  } */

function actualizardep(paisSeleccionado) {
  var departamento = document.getElementById("departamento");
  departamento.innerHTML = ""; // Limpiar la lista de ciudades
  var municipio = document.getElementById("municipio");
  municipio.innerHTML = ""; // Limpiar la lista de ciudades
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var departamentos= this.responseText;
          departamento.innerHTML = departamentos; // Actualizar la lista de ciudades
      }
  };
  xhr.open("POST", "/obtenerdepartamentog", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("pais=" + paisSeleccionado);
}

function actualizarmun(depSeleccionado) {

  var municipio = document.getElementById("municipio");
  municipio.innerHTML = ""; // Limpiar la lista de ciudades
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var municipios= this.responseText;
          municipio.innerHTML = municipios; // Actualizar la lista de ciudades
      }
  };
  xhr.open("POST", "/obtenermunicipiog", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("departamento=" + depSeleccionado);
}

function contribuyente(){
  var select = document.getElementById("tipoPersona");
  var hiddenDiv = document.getElementById("ocultarContribuyente");
  var input4 = document.getElementById("tipoContribuyente");
  if(select.value==="1"){    
    hiddenDiv.classList.add("d-none");
    input4.required = false;
    } else {
    hiddenDiv.classList.remove("d-none");
    input4.required = true;
      }
    }

    function formatNIT(input) {
      // Eliminar cualquier caracter que no sea un número
      let value = input.value.replace(/\D/g, '');
    
      // Formatear basado en la longitud del valor
      if (value.length == 9) {
        value = value.slice(0, 8) + '-' + value.slice(8);
      } else if (value.length >= 14) {
        value = value.padStart(14, '0');
        value = value.slice(0, 4) + '-' + value.slice(4, 10) + '-' + value.slice(10, 13) + '-' + value.slice(13);
      }
    
      // Actualizar el valor del input
      input.value = value;
    }

    function formatIdentidad(input) {
      // Eliminar cualquier caracter que no sea un número
      let value = input.value.replace(/\D/g, '');
    
      // Formatear basado en la longitud del valor
      if (value.length == 9) {
        value = value.slice(0, 8) + '-' + value.slice(8);
      }
    
      // Actualizar el valor del input
      input.value = value;
    }

    function Eleccionmetodo(){
      var select = document.getElementById("pago");
      var input1 = document.getElementById("banco");
      var input2 = document.getElementById("tipocuenta");
      var input3 = document.getElementById("cuenta");
      var input4 = document.getElementById("capturacuenta");
      if (select.value ==='1'){
        document.getElementById('divmetodo').classList.remove('d-none');
        input1.required = true;
        input2.required = true;
        input3.required = true;
        input4.required = true;
      } else {
        document.getElementById('divmetodo').classList.add('d-none');
        input1.required = false;
        input2.required = false;
        input3.required = false;
        input4.required = false;
      }
    }

    // Selecciona todos los inputs de tipo archivo
    var fileInputs = document.querySelectorAll('input[type="file"]');
    
    fileInputs.forEach(function(input) {
      input.addEventListener('change', function(e) {
        var file = e.target.files[0]; // Obtiene el archivo seleccionado
    
        // Verifica si el archivo es una imagen
        if (file && !file.type.startsWith('image/')) {
          // Usa Swal.fire para mostrar el mensaje de error
          Swal.fire({
            icon: 'error',
            title: 'Archivo no válido',
            text: 'Por favor, seleccione un archivo de imagen válido.',
            confirmButtonText: 'Aceptar'
          }).then((result) => {
            if (result.value) {
              // Limpia el campo de selección de archivo
              e.target.value = '';
            }
          });
        }
      });
    });

    function guardar(){
      var form = document.getElementById("formualrio");
      event.preventDefault();
      if (!form.checkValidity()) {
          event.preventDefault();
          event.stopPropagation();
          Swal.fire({
              icon: 'error',
              title: 'Oops...',
              text: 'Por favor, completa el formulario correctamente antes de enviar.',
          });
          return;
      }
      var formData = new FormData(form);
      var xhr = new XMLHttpRequest();
      xhr.onreadystatechange = function() {
          if (this.readyState == 4) {
              Swal.close(); // Cierra el Swal de "Subiendo..." independientemente del estado de la respuesta
              if (this.status == 200) {
                  var respuesta = this.responseText;
                  if (respuesta == "1") {
                      Swal.fire({
                          title: 'Respuesta correctamente',
                          icon: 'success'
                      }).then((result) => {
                        location.reload();
                      });
                  } else {
                      Swal.fire({
                          icon: 'error',
                          title: 'Oops...',
                          text: 'Error al registrar participante' + respuesta,
                      });
                  }
              } else if (this.status == 500) {
                  // Manejo específico del estado 500
                  Swal.fire({
                      icon: 'error',
                      title: 'Error del servidor',
                      text: 'Ha ocurrido un error interno en el servidor.',
                  });
              } else {
                  // Manejo de otros estados de error
                  Swal.fire({
                      icon: 'error',
                      title: 'Error de conexión',
                      text: 'No se pudo establecer conexión con el servidor. Estado: ' + this.status,
                  });
              }
          }
      };
      xhr.open("POST", "/finanzas/datos/post", true);
      Swal.fire({
          title: 'Subiendo...',
          allowOutsideClick: false,
          showConfirmButton: false,
          didOpen: () => {
              Swal.showLoading()
          },
      });
      xhr.send(formData);
  }