// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
  'use strict'

  // Fetch all the forms we want to apply custom Bootstrap validation styles to
  const forms = document.querySelectorAll('.needs-validation')

  // Loop over them and prevent submission
  Array.from(forms).forEach(form => {
    form.addEventListener('submit', event => {
      if (!form.checkValidity()) {
        event.preventDefault()
        event.stopPropagation()
      }
    }, false)
  })
})()

  function validateFormat() {
    var formatInput = document.getElementById("DUIM").value;
    var formatRegex = /^\d{1}-\d{3}-\d{3,4}$/;
       var select = document.getElementById("nacionalidad");
      if(select.value==="1"||select.value===""){
            if (!formatRegex.test(formatInput)) {
          var alertContainer = document.getElementById("alertContainer");
          alertContainer.innerHTML = '<div class="alert alert-danger" role="alert">El formato del documento de identidad no es válido.</div>';
          alertContainer.style.display = "block";
          return false;
    }
    return true;
  }
  }
 
  function hideAlert() {
    var alertContainer = document.getElementById("alertContainer");
    alertContainer.style.display = "none";
  }
 
function activaselect51(){
var select = document.getElementById("nacionalidad");
var labelOutput = document.getElementById("DUIMl");
if(select.value==="1"){    
  labelOutput.innerHTML = "Escribe el número de tu documento de identidad:"
  // document.getElementById("DUIM").placeholder = "0-000-000 o 0-000-0000";
    } else {labelOutput.innerHTML = "Escribe el número de tu documento de identidad"
    document.getElementById("DUIM").placeholder = ""
  } 
  } 
  
      function activaselect54(){
        var select = document.getElementById("nacionalidad2");
        if(select.value==="1"){    
          document.getElementById("DUIP").placeholder = "0-000-000 o 0-000-0000";
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

          const miFormulario = document.querySelector('#formualrioRegistro'); 
function quitar(id){
          var pregunta = document.getElementById(id);
          pregunta.classList.remove('pregunta-vacia');
        }
          
          function prueba() {
            
            var pregunta = document.getElementById("pregunta");
            var pregunta2 = document.getElementById("pregunta2");
            var pregunta3 = document.getElementById("pregunta3");
            var pregunta4 = document.getElementById("pregunta4");
            var pregunta5 = document.getElementById("pregunta5");
            var respuestaSeleccionada = miFormulario.querySelector('input[name="p1"]:checked');
            var respuestaSeleccionada2 = miFormulario.querySelector('input[name="p2"]:checked');
            var respuestaSeleccionada3= miFormulario.querySelector('input[name="p3"]:checked');
            var respuestaSeleccionada4 = miFormulario.querySelector('input[name="p4"]:checked');
            var respuestaSeleccionada5 = miFormulario.querySelector('input[name="p5"]:checked');
            
            if (!respuestaSeleccionada) {
              pregunta.classList.add('pregunta-vacia');
            } 
            if (!respuestaSeleccionada2) {
              pregunta2.classList.add('pregunta-vacia');
            } 
            if (!respuestaSeleccionada3) {
              pregunta3.classList.add('pregunta-vacia');
            } 
            if (!respuestaSeleccionada4) {
              pregunta4.classList.add('pregunta-vacia');
            } 
            if (!respuestaSeleccionada5) {
              pregunta5.classList.add('pregunta-vacia');
            } 
            
          }        