// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
    'use strict'
  
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation')
  
    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
      form.addEventListener('submit', event => {
        console.log('Formulario enviado, validando...');
        if (!form.checkValidity()) {
          event.preventDefault()
          event.stopPropagation()
        }
      }, false)
    })
  })()

  function mostrarPreguntaOtro(checkbox){
    var hiddenDiv1 = document.getElementById("preguntaOtro");
    var input = document.getElementById("otro");
    var errorMensaje = document.getElementById("errorMensaje");

    // Comprobar si el checkbox con valor 4 está seleccionado
    var isChecked = document.querySelector('input[name="razon1"][value="4"]').checked;

    // Comprobar si al menos un checkbox está seleccionado
    var anyChecked = document.querySelectorAll('input[name="razon1"]:checked').length > 0;

    if (isChecked) {
      hiddenDiv1.classList.remove("d-none");
      input.required = true;
    } else {
      hiddenDiv1.classList.add("d-none");
      input.required = false;
      input.value = "";
    }

  // Mostrar u ocultar el mensaje de error
    if (anyChecked) {
        errorMensaje.style.display = "none"; // Ocultar mensaje de error
    } else {
        errorMensaje.style.display = "block"; // Mostrar mensaje de error
    }
  }

    function activaselect4(){
    var select = document.getElementById("FechaN");
    var hiddenDiv = document.getElementById("ocultar");
    var birthDate = new Date(select.value);
    var diff_ms = Date.now() -birthDate ;
    var age_dt = new Date(diff_ms); 
    var age= Math.abs(age_dt.getUTCFullYear() - 1970);
            hiddenDiv.classList.remove("d-none");
    var input1= document.getElementById("sexo");
    var input2= document.getElementsByName("razon1");
    var input3= document.getElementsByName("satis1");
    var input4= document.getElementsByName("util1");
    var input5= document.getElementsByName("considerar1");
    var input6= document.getElementsByName("recomendar");
        
        if(age>=18){    
        hiddenDiv.classList.remove("d-none");
    
            input1.required = true;
            input2.required = true;
            input3.required = true;
            input4.required = true;
            input5.required = true;
            input6.required = true;
        }
        else{
            var alertContainer = document.getElementById("alertContainer2");
            alertContainer.innerHTML = '<div class="alert alert-danger" role="alert">Parece que has escrito mal tu fecha de nacimiento, por favor revisar.</div>';
            alertContainer.style.display = "block";
        
            hiddenDiv.classList.add("d-none");
            input1.required = false;
            input2.required = false;
            input3.required = false;
            input4.required = false;
            input5.required = false;
            input6.required = false;
        }
        }

    function hideAlert2() {
    var alertContainer = document.getElementById("alertContainer2");
    alertContainer.style.display = "none";
    }

    document.getElementById('formualrioRegistro').addEventListener('submit', function(event) {
      var checkboxes = document.querySelectorAll('input[name="razon1"]');
      var alMenosUnoMarcado = false;
  
      for (var i = 0; i < checkboxes.length; i++) {
          if (checkboxes[i].checked) {
              alMenosUnoMarcado = true;
              break;
          }
      }
  
      if (!alMenosUnoMarcado) {
        event.preventDefault(); // Evitar el envío del formulario
        var errorMensaje = document.getElementById('errorMensaje');
        errorMensaje.style.display = 'block'; // Mostrar mensaje de error
        errorMensaje.focus(); // Hacer foco en el mensaje de error
      } else {
        document.getElementById('errorMensaje').style.display = 'none'; // Ocultar mensaje de error si está marcado
      }
    });