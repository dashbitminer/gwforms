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

  function activaselect1() {
    var radios = document.querySelectorAll("input[name='gruporadio']");
    var hiddenDiv = document.getElementById("otrarespuesta");
    var input = document.getElementById("otrar");
    
    var mostrarDiv = false;
    var hacerRequired = false;
  
    radios.forEach(function(radio) {
      if (radio.value === "3" && radio.checked) {
        mostrarDiv = true;
        hacerRequired = true;
      }
    });
  
    if (mostrarDiv) {
      hiddenDiv.classList.remove("d-none");
    } else {
      hiddenDiv.classList.add("d-none");
    }
  
    input.required = hacerRequired;
  }

  function validarRadioButtons1() {
    var radioButtons = document.querySelectorAll("input[name='funcional']");
    var radioSeleccionado = false;
  
    // Verifica si al menos uno está seleccionado
    radioButtons.forEach(function(radio) {
      if (radio.checked) {
        radioSeleccionado = true;
        return;
      }
    });
  
    if (!radioSeleccionado) {
      alert("Selecciona al menos una opción (1 ,..., 10) para saber el nivel de funcionalidad de la herramienta.");
      return false; // Evita el envío del formulario
    }
  
    return true; // Permite el envío del formulario si al menos uno está seleccionado
  }

  function validarRadioButtons2() {
    var radioButtons = document.querySelectorAll("input[name='confianza']");
    var radioSeleccionado = false;
  
    // Verifica si al menos uno está seleccionado
    radioButtons.forEach(function(radio) {
      if (radio.checked) {
        radioSeleccionado = true;
        return;
      }
    });
  
    if (!radioSeleccionado) {
      alert("Selecciona al menos una opción (1 ,..., 10) para saber el nivel de confianza que tenias antes de usar la herramienta.");
      return false; // Evita el envío del formulario
    }
  
    return true; // Permite el envío del formulario si al menos uno está seleccionado
  }