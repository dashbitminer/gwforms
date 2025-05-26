
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
      else if (document.getElementById("alertContainer").style.display === "block") {
        event.preventDefault();
        alert("No se puede enviar el formulario mientras este incorrecto el DUI");
      }
    }, false)
  })
})()

function activaselect4() {
  var select = document.getElementById("FechaN");
  var hiddenDiv3 = document.getElementById("ocultar");

  var birthDate = new Date(select.value);
  var diff_ms = Date.now() - birthDate;
  var age_dt = new Date(diff_ms);
  var age = Math.abs(age_dt.getUTCFullYear() - 1970);
  var output1 = document.getElementById("edadl")
  output1.innerHTML = age;

  if (age >= 18) {
    hiddenDiv3.classList.remove("d-none");

  } else {
    hiddenDiv3.classList.add("d-none");
  }

}

function activa1() {
  var input = document.getElementById("nombre");
  var input2 = document.getElementById("apellido");
  var output2 = document.getElementById("nombrep2");
  output2.innerHTML = input.value + " " + input2.value;

}

document.getElementById('identidad').addEventListener('input', function (e) {
  this.value = this.value.replace(/[^A-Za-z0-9]/g, '');
});


