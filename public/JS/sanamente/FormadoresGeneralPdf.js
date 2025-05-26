// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
    'use strict'
  
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation')
  
    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
      form.addEventListener('submit', event => {
        validateFormat()
        if (!form.checkValidity()) {
          event.preventDefault()
          event.stopPropagation()
        }
      }, false)
    })
  })()
  
            function actualizarmun(paisSeleccionado) {
              var ciudadSelect = document.getElementById("departamento2");
              ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
              var sede = document.getElementById("sede2");
              sede.innerHTML = ""; // Limpiar la lista de ciudades
              // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
              var xhr = new XMLHttpRequest();
              xhr.onreadystatechange = function() {
                  if (this.readyState == 4 && this.status == 200) {
                      var pcionesmunsede= this.responseText;
                      ciudadSelect.innerHTML = pcionesmunsede; // Actualizar la lista de ciudades
                  }
              };
              xhr.open("POST", "/obtenermunsedesanamente", true);
              xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
              xhr.send("depsede=" + paisSeleccionado);
            }
            
            function actualizarsede(paisSeleccionado) {
              var ciudadSelect = document.getElementById("sede2");
              ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
              // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
              var xhr = new XMLHttpRequest();
              xhr.onreadystatechange = function() {
                  if (this.readyState == 4 && this.status == 200) {
                      var opcionesmunsede= this.responseText;
                      ciudadSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
                  }
              };
              xhr.open("POST", "/obtenersedesanamente", true);
              xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
              xhr.send("munsede=" + paisSeleccionado);
            }
  
            function mostrarPregunta1() {
              var opcion1 = document.getElementById("p1");
              var opcion2 = document.getElementById("p2");
              var opcion3 = document.getElementById("p3");
              var opcion4 = document.getElementById("p4");
                if (opcion1.checked||opcion2.checked||opcion3.checked||opcion4.checked) {
                  document.getElementById("pregunta").classList.remove("d-none")
                } else {
                  document.getElementById("pregunta").classList.add("d-none");
                }
              }
  
  
            function mostrarPregunta2() {
              var opcion1 = document.getElementById("p9");
              var opcion2 = document.getElementById("p10");
           
                if (opcion1.checked||opcion2.checked) {
                    document.getElementById("pregunta2").classList.remove("d-none")
                  } else {
                    document.getElementById("pregunta2").classList.add("d-none");
                  }
              }
  
            function mostrarPregunta3() {
              var opcion1 = document.getElementById("p17");
              var opcion2 = document.getElementById("p18");
              var opcion3 = document.getElementById("p19");
              var opcion4 = document.getElementById("p20");
              var opcion5 = document.getElementById("p21");
              var opcion6 = document.getElementById("p22");
       
                if (opcion1.checked||opcion2.checked||opcion3.checked||opcion4.checked||opcion5.checked||opcion6.checked) {
                  document.getElementById("pregunta3").classList.remove("d-none")
                } else {
                  document.getElementById("pregunta3").classList.add("d-none");
                }
              }
  
              function mostrarPregunta4() {
                var opcion1 = document.getElementById("p30");
                var opcion2 = document.getElementById("p31");
                var opcion3 = document.getElementById("p32");
                var opcion4 = document.getElementById("p33");
                var opcion5 = document.getElementById("p34");
  
                  if (opcion1.checked||opcion2.checked||opcion3.checked||opcion4.checked||opcion5.checked) {
                    document.getElementById("pregunta4").classList.remove("d-none")
                  } else {
                    document.getElementById("pregunta4").classList.add("d-none");
                  }
                }
  
                function mostrarPregunta5() {
                  var opcion1 = document.getElementById("p40");
                  var opcion2 = document.getElementById("p41");
                  var opcion3 = document.getElementById("p42");              
    
                    if (opcion1.checked||opcion2.checked||opcion3.checked) {
                      document.getElementById("pregunta5").classList.remove("d-none")
                    } else {
                      document.getElementById("pregunta5").classList.add("d-none");
                    }
                  }
  
                  function mostrarPregunta6() {
                    var opcion1 = document.getElementById("p50");
                    var opcion2 = document.getElementById("p51");
                    var opcion3 = document.getElementById("p52");
                    var opcion4 = document.getElementById("p53");
      
                      if (opcion1.checked||opcion2.checked||opcion3.checked||opcion4.checked) {
                        document.getElementById("pregunta6").classList.remove("d-none")
                      } else {
                        document.getElementById("pregunta6").classList.add("d-none");
                      }
                    }
  
                    function mostrarPregunta7() {
                      var opcion1 = document.getElementById("p59");
                      var opcion2 = document.getElementById("p60");
                      var opcion3 = document.getElementById("p61");
                      var opcion4 = document.getElementById("p62");
        
                        if (opcion1.checked||opcion2.checked||opcion3.checked||opcion4.checked) {
                          document.getElementById("pregunta7").classList.remove("d-none")
                        } else {
                          document.getElementById("pregunta7").classList.add("d-none");
                        }
                      }
  
  // Convert HTML content to PDF
  // Convert HTML content to PDF
  function Convert_HTML_To_PDF() {
    
    var form = document.getElementById("formualrioRegistro");
    var invalidElements = form.querySelectorAll('.is-invalid');
    event.preventDefault();
    if (!form.checkValidity() || invalidElements.length > 0) {
      event.preventDefault();
      event.stopPropagation();
      Swal.fire({
        icon: 'error',
        title: 'Oops...',
        text: 'Por favor, completa el formulario correctamente antes de enviar.',
      });
      return;
    }
    document.getElementById("boton").setAttribute("disabled", "true");
  
    event.preventDefault();
    window.jsPDF = window.jspdf.jsPDF;
    // Get the HTML content to be printed
    var elementHTML = document.querySelector("#contentToPrint");
  
    var doc = new jsPDF();
  
    // Convert the HTML content to canvas with reduced scale
    html2canvas(elementHTML, { scale: 1 }).then(function (canvas) {
  
      // Add the canvas to the first page of the PDF document with reduced quality
      var imgData = canvas.toDataURL('image/jpeg');
      doc.addImage(imgData, 'JPEG', 10, 10, 190, 269);
  
      // Add a new page to the PDF document
      doc.addPage();
  
      // Capture another HTML element as an image
      var elementToCapture2 = document.querySelector("#contentToPrint2");
      html2canvas(elementToCapture2, { scale: 1 }).then(function (canvas) {
  
        // Add the captured image to the new page of the PDF document
        var imgData = canvas.toDataURL('image/jpeg');
        doc.addImage(imgData, 'JPEG', 10, 10, 190, 269);
        doc.addPage();
  
        // Capture a third HTML element as an image
        var elementToCapture3 = document.querySelector("#contentToPrint3");
        html2canvas(elementToCapture3, { scale: 1 }).then(function (canvas) {
  
          // Add the captured image to the new page of the PDF document
          var imgData = canvas.toDataURL('image/jpeg');
          doc.addImage(imgData, 'JPEG', 10, 10, 190, 269);
          doc.addPage();
  
          // Capture a fourth HTML element as an image
          var elementToCapture4 = document.querySelector("#contentToPrint4");
          html2canvas(elementToCapture4, { scale: 1 }).then(function (canvas) {
  
            // Add the captured image to the new page of the PDF document
            var imgData = canvas.toDataURL('image/jpeg');
            doc.addImage(imgData, 'JPEG', 10, 10, 190, 269);
  
            // Save the PDF document
            doc.save('documento.pdf');
          });
        });
      });
    });
  }

  function activaselectform(){
    var select = document.getElementById("formaciones");
    var hiddenDiv1 = document.getElementById("hidden-question-div-fentrenamiento");
    var hiddenDiv2 = document.getElementById("hidden-question-div-modulos");
    var input1 = document.getElementById("fentrenamiento");
    var input2 = document.getElementById("m1");
    var input3 = document.getElementById("m2");
    var input4 = document.getElementById("m3");
    if(select.value==="43"){    
    hiddenDiv1.classList.remove("d-none");
    hiddenDiv2.classList.add("d-none");
    input1.required = true;
    input2.checked = false;
    input3.checked = false;
    input4.checked = false;
    } else if (select.value==="106"){
    hiddenDiv2.classList.remove("d-none");
    hiddenDiv1.classList.add("d-none");
    input1.required = false;
    input1.value = "";
    } else {
    hiddenDiv1.classList.add("d-none");
    hiddenDiv2.classList.add("d-none");
    input1.required = false;
    input1.value = "";
    input2.checked = false;
    input3.checked = false;
    input4.checked = false;
    }
    }