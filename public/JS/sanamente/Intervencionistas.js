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
              var opcion1 = document.getElementById("p67");
              var opcion2 = document.getElementById("p68");
              var opcion3 = document.getElementById("p69");
              var opcion4 = document.getElementById("p70");
              var opcion5 = document.getElementById("p71");
                if (opcion1.checked||opcion2.checked||opcion3.checked||opcion4.checked||opcion5.checked) {
                  document.getElementById("pregunta").classList.remove("d-none")
                } else {
                  document.getElementById("pregunta").classList.add("d-none");
                }
              }
  
  
            function mostrarPregunta2() {
              var opcion1 = document.getElementById("p77");
              var opcion2 = document.getElementById("p78");
              var opcion3 = document.getElementById("p79");
              var opcion4 = document.getElementById("p80");
           
                if (opcion1.checked||opcion2.checked||opcion3.checked||opcion4.checked) {
                    document.getElementById("pregunta2").classList.remove("d-none")
                  } else {
                    document.getElementById("pregunta2").classList.add("d-none");
                  }
              }
  
            function mostrarPregunta3() {
              var opcion1 = document.getElementById("p88");
              var opcion2 = document.getElementById("p89");
              var opcion3 = document.getElementById("p90");
              var opcion4 = document.getElementById("p91");
              var opcion5 = document.getElementById("p92");
              var opcion6 = document.getElementById("p93");
       
                if (opcion1.checked||opcion2.checked||opcion3.checked||opcion4.checked||opcion5.checked||opcion6.checked) {
                  document.getElementById("pregunta3").classList.remove("d-none")
                } else {
                  document.getElementById("pregunta3").classList.add("d-none");
                }
              }
  
              function mostrarPregunta4() {
                var opcion1 = document.getElementById("p100");
                var opcion2 = document.getElementById("p101");
                var opcion3 = document.getElementById("p102");
                var opcion4 = document.getElementById("p103");
  
                  if (opcion1.checked||opcion2.checked||opcion3.checked||opcion4.checked) {
                    document.getElementById("pregunta4").classList.remove("d-none")
                  } else {
                    document.getElementById("pregunta4").classList.add("d-none");
                  }
                }
  
                function mostrarPregunta5() {
                  var opcion1 = document.getElementById("p108");
                  var opcion2 = document.getElementById("p109");
                  var opcion3 = document.getElementById("p110");
                  var opcion4 = document.getElementById("p111");
                  var opcion5 = document.getElementById("p112");                
    
                    if (opcion1.checked||opcion2.checked||opcion3.checked||opcion4.checked||opcion5.checked) {
                      document.getElementById("pregunta5").classList.remove("d-none")
                    } else {
                      document.getElementById("pregunta5").classList.add("d-none");
                    }
                  }
  
                  function mostrarPregunta6() {
                    var opcion1 = document.getElementById("p117");
                    var opcion2 = document.getElementById("p118");
                    var opcion3 = document.getElementById("p119");
                    var opcion4 = document.getElementById("p120");
                    var opcion5 = document.getElementById("p121");
      
                      if (opcion1.checked||opcion2.checked||opcion3.checked||opcion4.checked||opcion5.checked) {
                        document.getElementById("pregunta6").classList.remove("d-none")
                      } else {
                        document.getElementById("pregunta6").classList.add("d-none");
                      }
                    }
  
                    function mostrarPregunta7() {
                      var opcion1 = document.getElementById("p127");
                      var opcion2 = document.getElementById("p128");
                      var opcion3 = document.getElementById("p129");
        
                        if (opcion1.checked||opcion2.checked||opcion3.checked) {
                          document.getElementById("pregunta7").classList.remove("d-none")
                        } else {
                          document.getElementById("pregunta7").classList.add("d-none");
                        }
                      }
  
  // Convert HTML content to PDF
  function Convert_HTML_To_PDF() {
  
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
        doc.addImage(imgData, 'JPEG', 10, 10,190, 269);
        doc.addPage();
  
        // Capture a third HTML element as an image
        var elementToCapture3 = document.querySelector("#contentToPrint3");
        html2canvas(elementToCapture3, { scale: 1 }).then(function (canvas) {
  
          // Add the captured image to the new page of the PDF document
          var imgData = canvas.toDataURL('image/jpeg');
          doc.addImage(imgData, 'JPEG', 10, 10,190, 269);
  
          // Add the captured image to the new page of the PDF document
          var imgData = canvas.toDataURL('image/jpeg');
          doc.addImage(imgData, 'JPEG', 10, 10,190, 269);
          doc.addPage();
    
          // Capture a third HTML element as an image
          var elementToCapture4 = document.querySelector("#contentToPrint4");
          html2canvas(elementToCapture4, { scale: 1 }).then(function (canvas) {
    
            // Add the captured image to the new page of the PDF document
            var imgData = canvas.toDataURL('image/jpeg');
            doc.addImage(imgData, 'JPEG', 10, 10,190, 269);
    
          // Get the raw PDF data as a byte array
  
          pdfData = doc.output('arraybuffer');
                  var formData = new FormData();
                  formData.append('pdf', new Blob([pdfData], { type: 'application/pdf' }))
                  const email = document.getElementById('email').value;
                  formData.append('email', email);
          // Send the PDF data to the server using an AJAX request
         fetch('/sanamente/equip/email', {
    method: 'POST',
    body: formData
  })
          .then(function(response) {
            // Handle the server response here
            document.getElementById("formualrioRegistro").submit();
          })
          .catch(function(error) {
            console.error('Error sending PDF data to server:', error);
          });
        });
        });
      });
    });
  }