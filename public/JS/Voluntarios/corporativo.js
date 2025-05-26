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

  function eventos(paisSeleccionado) {
    var ciudadSelect = document.getElementById("evento");
    ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades

    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesCiudades = this.responseText;
            ciudadSelect.innerHTML = opcionesCiudades; // Actualizar la lista de ciudades
        }
    };
    xhr.open("POST", "/obtenereventos", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("pais=" + paisSeleccionado);
  }
  function eventosgps() {

    var ciudadSelect = document.getElementById("evento");
    var hiddenDiv = document.getElementById("mx");
    var hiddenDiv2 = document.getElementById("pn");
    var hiddenDiv3 = document.getElementById("cr");
    var hiddenDiv4 = document.getElementById("gt");

  if(ciudadSelect.value==="1"){
    hiddenDiv.classList.add("d-none");
    hiddenDiv2.classList.add("d-none");
    hiddenDiv3.classList.add("d-none");
    hiddenDiv4.classList.remove("d-none");
  }

  if(ciudadSelect.value==="2"){
    hiddenDiv.classList.add("d-none");
    hiddenDiv2.classList.remove("d-none");
    hiddenDiv3.classList.add("d-none");
    hiddenDiv4.classList.add("d-none");
  }

  if(ciudadSelect.value==="3"){
    hiddenDiv.classList.remove("d-none");
    hiddenDiv2.classList.add("d-none");
    hiddenDiv3.classList.add("d-none");
    hiddenDiv4.classList.add("d-none");
  }

  if(ciudadSelect.value==="4"){
    hiddenDiv.classList.add("d-none");
    hiddenDiv2.classList.add("d-none");
    hiddenDiv3.classList.remove("d-none");
    hiddenDiv4.classList.add("d-none");
  }
}

function activa1(){
var input = document.getElementById("nombreCompleto");
var output = document.getElementById("nombrep2");
  output.innerHTML = input.value;

}

function activaselect4(){
  var select = document.getElementById("FechaN");
  var birthDate = new Date(select.value);
  var diff_ms = Date.now() -birthDate ;
  var age_dt = new Date(diff_ms); 
  var age= Math.abs(age_dt.getUTCFullYear() - 1970);
  var output1 = document.getElementById("edadl");
    output1.innerHTML = age;
}
