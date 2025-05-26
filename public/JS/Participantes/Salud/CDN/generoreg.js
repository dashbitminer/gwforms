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

function ocultarppal(selectedRadio){
  var div= document.getElementById("ocultar");

  if (selectedRadio.value == "1") {
    div.classList.remove("d-none");
  }else if (selectedRadio.value == "0") {
    div.classList.add("d-none");
  }
}

function actualizardepto(paisSeleccionado) {
  var ciudadSelect = document.getElementById("pais2");
  ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var opcionesmunsede= this.responseText;
          ciudadSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
        }
  };
  xhr.open("POST", "/obtenerdepsedesalud", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("paissede=" + paisSeleccionado);
  ciudadSelect.required = true;
}

function actualizarmun(depSeleccionado) {
  var ciudadSelect = document.getElementById("departamento2");
  ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var pcionesmunsede= this.responseText;
          ciudadSelect.innerHTML = pcionesmunsede; // Actualizar la lista de ciudades
      }
  };
  xhr.open("POST", "/obtenermunsedesalud", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("depsede=" + depSeleccionado);
}

function actualizarsede(munSeleccionado) {
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
  xhr.open("POST", "/obtenersedesalud", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("munsede=" + munSeleccionado);
}

function ocultardiv1(selectedRadio){
  var div= document.getElementById("pregunta6");

  if (selectedRadio.value == "1" ||selectedRadio.value == "2" ) {
    div.classList.remove("d-none");
  }else{
    div.classList.add("d-none");}
}

function ocultardiv2(selectedRadio){
  var div= document.getElementById("pregunta10");

  if (selectedRadio.value == "1" ) {
    div.classList.remove("d-none");
  }else{
    div.classList.add("d-none");}
}

function ocultardiv3(selectedCheck){
  var div= document.getElementById("pregunta7");

  if (selectedCheck.checked) {
    div.classList.remove("d-none");
  }else{
    div.classList.add("d-none");}
}

 function controlCheckboxes(selectedRadio) {
  const checkbox1 = document.getElementById('sr1-opc1');
  const checkbox2 = document.getElementById('sr1-opc2');
  const checkbox3 = document.getElementById('sr1-opc3');

  if (selectedRadio.value == "1") {
    checkbox1.disabled = false;
    checkbox2.disabled = false;
    checkbox3.disabled = false;
  } else if (selectedRadio.value == "0") {
    checkbox1.disabled = true;
    checkbox2.disabled = true;
    checkbox3.disabled = true;

    // Deseleccionar checkboxes si están deshabilitados
    checkbox1.checked = false;
    checkbox2.checked = false;
    checkbox3.checked = false;
  }
}

function controlCheckboxes2(selectedRadio) {
  const checkbox1 = document.getElementById('sr2-opc1');
  const checkbox2 = document.getElementById('sr2-opc2');
  const checkbox3 = document.getElementById('sr2-opc3');

  if (selectedRadio.value == "1") {
    checkbox1.disabled = false;
    checkbox2.disabled = false;
    checkbox3.disabled = false;
  } else if (selectedRadio.value == "0") {
    checkbox1.disabled = true;
    checkbox2.disabled = true;
    checkbox3.disabled = true;

    // Deseleccionar checkboxes si están deshabilitados
    checkbox1.checked = false;
    checkbox2.checked = false;
    checkbox3.checked = false;
  }
}
function controlCheckboxes3(selectedRadio) {
  const checkbox1 = document.getElementById('sr3-opc1');
  const checkbox2 = document.getElementById('sr3-opc2');
  const checkbox3 = document.getElementById('sr3-opc3');

  if (selectedRadio.value == "1") {
    checkbox1.disabled = false;
    checkbox2.disabled = false;
    checkbox3.disabled = false;
  } else if (selectedRadio.value == "0") {
    checkbox1.disabled = true;
    checkbox2.disabled = true;
    checkbox3.disabled = true;

    // Deseleccionar checkboxes si están deshabilitados
    checkbox1.checked = false;
    checkbox2.checked = false;
    checkbox3.checked = false;
  }
}
function controlCheckboxes4(selectedRadio) {
  const checkbox1 = document.getElementById('sr4-opc1');
  const checkbox2 = document.getElementById('sr4-opc2');
  const checkbox3 = document.getElementById('sr4-opc3');

  if (selectedRadio.value == "1") {
    checkbox1.disabled = false;
    checkbox2.disabled = false;
    checkbox3.disabled = false;
  } else if (selectedRadio.value == "0") {
    checkbox1.disabled = true;
    checkbox2.disabled = true;
    checkbox3.disabled = true;

    // Deseleccionar checkboxes si están deshabilitados
    checkbox1.checked = false;
    checkbox2.checked = false;
    checkbox3.checked = false;
  }
}
function controlCheckboxes5(selectedRadio) {
  const checkbox1 = document.getElementById('sr5-opc1');
  const checkbox2 = document.getElementById('sr5-opc2');
  const checkbox3 = document.getElementById('sr5-opc3');

  if (selectedRadio.value == "1") {
    checkbox1.disabled = false;
    checkbox2.disabled = false;
    checkbox3.disabled = false;
  } else if (selectedRadio.value == "0") {
    checkbox1.disabled = true;
    checkbox2.disabled = true;
    checkbox3.disabled = true;

    // Deseleccionar checkboxes si están deshabilitados
    checkbox1.checked = false;
    checkbox2.checked = false;
    checkbox3.checked = false;
  }
}

function controlCheckboxesSV(selectedRadio) {
  const checkbox1 = document.getElementById('sr1-opc1');
  const checkbox2 = document.getElementById('sr1-opc2');
  const checkbox3 = document.getElementById('sr1-opc3');
  const checkbox4 = document.getElementById('sr1-opc4');

  if (selectedRadio.value == "1") {
    checkbox1.disabled = false;
    checkbox2.disabled = false;
    checkbox3.disabled = false;
    checkbox4.disabled = false;
  } else if (selectedRadio.value == "0") {
    checkbox1.disabled = true;
    checkbox2.disabled = true;
    checkbox3.disabled = true;
    checkbox4.disabled = true;

    // Deseleccionar checkboxes si están deshabilitados
    checkbox1.checked = false;
    checkbox2.checked = false;
    checkbox3.checked = false;
    checkbox4.checked = false;
  }
}

function controlCheckboxes2SV(selectedRadio) {
  const checkbox1 = document.getElementById('sr2-opc1');
  const checkbox2 = document.getElementById('sr2-opc2');
  const checkbox3 = document.getElementById('sr2-opc3');
  const checkbox4 = document.getElementById('sr2-opc4');

  if (selectedRadio.value == "1") {
    checkbox1.disabled = false;
    checkbox2.disabled = false;
    checkbox3.disabled = false;
    checkbox4.disabled = false;
  } else if (selectedRadio.value == "0") {
    checkbox1.disabled = true;
    checkbox2.disabled = true;
    checkbox3.disabled = true;
    checkbox4.disabled = true;

    // Deseleccionar checkboxes si están deshabilitados
    checkbox1.checked = false;
    checkbox2.checked = false;
    checkbox3.checked = false;
    checkbox4.checked = false;
  }
}
function controlCheckboxes3SV(selectedRadio) {
  const checkbox1 = document.getElementById('sr3-opc1');
  const checkbox2 = document.getElementById('sr3-opc2');
  const checkbox3 = document.getElementById('sr3-opc3');
  const checkbox4 = document.getElementById('sr3-opc4');

  if (selectedRadio.value == "1") {
    checkbox1.disabled = false;
    checkbox2.disabled = false;
    checkbox3.disabled = false;
    checkbox4.disabled = false;
  } else if (selectedRadio.value == "0") {
    checkbox1.disabled = true;
    checkbox2.disabled = true;
    checkbox3.disabled = true;
    checkbox4.disabled = true;

    // Deseleccionar checkboxes si están deshabilitados
    checkbox1.checked = false;
    checkbox2.checked = false;
    checkbox3.checked = false;
    checkbox4.checked = false;
  }
}
function controlCheckboxes4SV(selectedRadio) {
  const checkbox1 = document.getElementById('sr4-opc1');
  const checkbox2 = document.getElementById('sr4-opc2');
  const checkbox3 = document.getElementById('sr4-opc3');
  const checkbox4 = document.getElementById('sr4-opc4');

  if (selectedRadio.value == "1") {
    checkbox1.disabled = false;
    checkbox2.disabled = false;
    checkbox3.disabled = false;
    checkbox4.disabled = false;
  } else if (selectedRadio.value == "0") {
    checkbox1.disabled = true;
    checkbox2.disabled = true;
    checkbox3.disabled = true;
    checkbox4.disabled = true;

    // Deseleccionar checkboxes si están deshabilitados
    checkbox1.checked = false;
    checkbox2.checked = false;
    checkbox3.checked = false;
    checkbox4.checked = false;
  }
}
function controlCheckboxes5SV(selectedRadio) {
  const checkbox1 = document.getElementById('sr5-opc1');
  const checkbox2 = document.getElementById('sr5-opc2');
  const checkbox3 = document.getElementById('sr5-opc3');
  const checkbox4 = document.getElementById('sr5-opc4');

  if (selectedRadio.value == "1") {
    checkbox1.disabled = false;
    checkbox2.disabled = false;
    checkbox3.disabled = false;
    checkbox4.disabled = false;
  } else if (selectedRadio.value == "0") {
    checkbox1.disabled = true;
    checkbox2.disabled = true;
    checkbox3.disabled = true;
    checkbox4.disabled = true;

    // Deseleccionar checkboxes si están deshabilitados
    checkbox1.checked = false;
    checkbox2.checked = false;
    checkbox3.checked = false;
    checkbox4.checked = false;
  }
}