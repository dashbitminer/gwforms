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
  
   
  function activaselect51(){
  var select = document.getElementById("nacionalidad");
  var labelOutput = document.getElementById("DUIMl");
  if(select.value==="1"){    
    labelOutput.innerHTML = "Escribe el número de tu Documento de Identidad:"
    document.getElementById("DUIM").placeholder = "0000000000000";
      } else {labelOutput.innerHTML = "Escribe el número de tu Documento de Identidad"
      document.getElementById("DUIM").placeholder = ""
    } 
    } 
    
        function activaselect54(){
          var select = document.getElementById("nacionalidad2");
          if(select.value==="1"){    
            document.getElementById("DUIP").placeholder = "0000000000000";
              } else {
              document.getElementById("DUIP").placeholder = ""
            } 
            }
  
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
             // $('#sede2').selectpicker()[0].selectpicker.destroy(); anterior
              $('#sede2').selectpicker('destroy');
              // $('#sede2').selectpicker('destroy');
              ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
              // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  
              var xhr = new XMLHttpRequest();
              xhr.onreadystatechange = function() {
                  if (this.readyState == 4 && this.status == 200) {
                    var opcionesmunsede= this.responseText;
                      ciudadSelect.innerHTML = opcionesmunsede;
                      $('#sede2').selectpicker(); // Actualizar la lista de ciudades
                  }
              };
              xhr.open("POST", "/obtenersedesanamente", true);
              xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
              xhr.send("munsede=" + paisSeleccionado);
            }