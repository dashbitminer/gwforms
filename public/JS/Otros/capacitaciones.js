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
  
        form.classList.add('was-validated')
        
      }, false)
  
    })
  })()
  
  function activaselect(){
  var select = document.getElementById("detallePago");
  var hiddenDiv = document.getElementById("nombreParticipanteD");
  var input = document.getElementById("nombreParticipante");
  if(select.value==="1"){    
   hiddenDiv.classList.remove("d-none");
    input.required = true;
      } else {
     hiddenDiv.classList.add("d-none");
      input.required = false;
    } 
    } 
  
  function activaselect2(){
  var select = document.getElementById("tipoContribuyente");
  var hiddenDiv = document.getElementById("numeroContactoD");
  var input = document.getElementById("numeroContacto");
  if(select.value==="1"){    
    hiddenDiv.classList.remove("d-none");
    input.required = true;
  
      } else {
     hiddenDiv.classList.add("d-none");
      input.required = false;
    } 
    } 
  