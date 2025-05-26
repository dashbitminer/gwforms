// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
    'use strict'
  
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation')
  
    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
      form.addEventListener('submit', event => {
        console.log('Formulario enviado, validando...');
        validateFormat()
        if (!form.checkValidity()) {
          event.preventDefault()
          event.stopPropagation()
        }
      }, false)
    })
  })()

  function mostrarPreguntaOtro(){
    var hiddenDiv1 = document.getElementById("preguntaOtro");
    var input1 = document.getElementById("tipo1");
    var input2 = document.getElementById("tipo2");
    var input3 = document.getElementById("tipo3");
    var input4 = document.getElementById("tipo4");
    var input5 = document.getElementById("tipo5");
    var input6 = document.getElementById("otro")
    if(input5.checked){    
    hiddenDiv1.classList.remove("d-none");
    input6.required = true;
    } else {
    hiddenDiv1.classList.add("d-none");
    input6.required = false;
    input6.value = "";
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
    var input2= document.getElementsByName("tipo1");
    var input3= document.getElementById("otro");
    var input4= document.getElementById("institucion");
    var input5= document.getElementById("cantServicio");
    var input6= document.getElementsByName("satis1");
    var input7= document.getElementsByName("op1");
    var input8= document.getElementById("sugerencias");
        
        if(age>=18){    
        hiddenDiv.classList.remove("d-none");
    
            input1.required = true;
            input2.required = true;
            input4.required = true;
            input5.required = true;
            input6.required = true;
            input7.required = true;
        }
        else{
            var alertContainer = document.getElementById("alertContainer2");
            alertContainer.innerHTML = '<div class="alert alert-danger" role="alert">Parece que has escrito mal tu fecha de nacimiento, por favor revisar.</div>';
            alertContainer.style.display = "block";
        
            hiddenDiv.classList.add("d-none");
            input1.required = false;
            input2.required = false;
            input4.required = false;
            input5.required = false;
            input6.required = false;
            input7.required = false;
        }
        }

    function hideAlert2() {
    var alertContainer = document.getElementById("alertContainer2");
    alertContainer.style.display = "none";
    }