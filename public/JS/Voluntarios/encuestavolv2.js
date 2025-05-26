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

function activaselect1(){
    var select = document.getElementById("p2");
    var hiddenDiv = document.getElementById("otrol");
    var input = document.getElementById("otro1")

    if(select.value==="13"){
      input.required=true;
      hiddenDiv.classList.remove("d-none");
    }
    else{
      input.required=false;
      input.value="";
      hiddenDiv.classList.add("d-none");
    }
}

/* function activaselect2() {
  var hiddenDiv = document.getElementById('otrol2');
  var input = document.getElementById('otro2');

  // Comprobar si la opción con valor 42 está seleccionada en el select múltiple
  var isSelected = document.querySelector('select[name="p3"] option[value="18"]:checked');

  if (isSelected) {
    hiddenDiv.classList.remove("d-none");
    input.required = true;
  } else {
    hiddenDiv.classList.add("d-none");
    input.required = false;
    input.value = ""; // Limpiar el valor cuando se oculta el div
  }
} */

function activaselect2(checkbox) {
  var hiddenDiv1 = document.getElementById('otrol2');
  var input = document.getElementById('otro2');
  
  // Comprobar si el checkbox con valor 6 está seleccionado
  var isChecked = document.querySelector('input[name="p3"][value="18"]').checked;
  
  if (isChecked) {
    hiddenDiv1.classList.remove("d-none");
    input.required = true;
  } else {
    hiddenDiv1.classList.add("d-none");
    input.required = false;
    input.value = ""; // Clear input value when hiding
  }
}

/* function activaselect3() {
  var hiddenDiv = document.getElementById('otrol3');
  var input = document.getElementById('otro3');

  // Comprobar si la opción con valor 42 está seleccionada en el select múltiple
  var isSelected = document.querySelector('select[name="p7"] option[value="42"]:checked');

  if (isSelected) {
    hiddenDiv.classList.remove("d-none");
    input.required = true;
  } else {
    hiddenDiv.classList.add("d-none");
    input.required = false;
    input.value = ""; // Limpiar el valor cuando se oculta el div
  }
} */

function activaselect3(checkbox) {
  var hiddenDiv1 = document.getElementById('otrol3');
  var input = document.getElementById('otro3');
  
  // Comprobar si el checkbox con valor 6 está seleccionado
  var isChecked = document.querySelector('input[name="p7"][value="42"]').checked;
  
  if (isChecked) {
    hiddenDiv1.classList.remove("d-none");
    input.required = true;
  } else {
    hiddenDiv1.classList.add("d-none");
    input.required = false;
    input.value = ""; // Clear input value when hiding
  }
}


function activaselect4(){
  var select = document.getElementById("p8");
  var hiddenDiv = document.getElementById("otrol4");
  var input = document.getElementById("otro4")

  if(select.value==="49"){
    input.required=true;
    hiddenDiv.classList.remove("d-none");
  }
  else{
    input.required=false;
    input.value="";
    hiddenDiv.classList.add("d-none");
  }
}

function activaselect5(){
  var select = document.getElementById("p28");
  var hiddenDiv = document.getElementById("otrol5");
  var input = document.getElementById("otro5")

  if(select.value==="86"){
    input.required=true;
    hiddenDiv.classList.remove("d-none");
  }
  else{
    input.required=false;
    input.value="";
    hiddenDiv.classList.add("d-none");
  }
}

/* function activaselect6() {
  var hiddenDiv = document.getElementById('otrol5');
  var input = document.getElementById('otro5');

  // Comprobar si la opción con valor 42 está seleccionada en el select múltiple
  var isSelected = document.querySelector('select[name="p28"] option[value="86"]:checked');

  if (isSelected) {
    hiddenDiv.classList.remove("d-none");
    input.required = true;
  } else {
    hiddenDiv.classList.add("d-none");
    input.required = false;
    input.value = ""; // Limpiar el valor cuando se oculta el div
  }
} */

function activaselect6(checkbox) {
  var hiddenDiv1 = document.getElementById('otrol5');
  var input = document.getElementById('otro5');
  
  // Comprobar si el checkbox con valor 6 está seleccionado
  var isChecked = document.querySelector('input[name="p28"][value="86"]').checked;
  
  if (isChecked) {
    hiddenDiv1.classList.remove("d-none");
    input.required = true;
  } else {
    hiddenDiv1.classList.add("d-none");
    input.required = false;
    input.value = ""; // Clear input value when hiding
  }
}

function activaselect7(){
  var select = document.getElementById("p9");
  var hiddenDiv = document.getElementById("razon1");
  var input = document.getElementById("xq5")

  if(select.value==="53"){
    input.required=true;
    hiddenDiv.classList.remove("d-none");
  }
  else{
    input.required=false;
    input.value="";
    hiddenDiv.classList.add("d-none");
  }
}

function activaselect8(){
  var select = document.getElementById("p23");
  var hiddenDiv = document.getElementById("razon2");
  var input = document.getElementById("xq6")

  if(select.value==="53"){
    input.required=true;
    hiddenDiv.classList.remove("d-none");
  }
  else{
    input.required=false;
    input.value="";
    hiddenDiv.classList.add("d-none");
  }
}

function activaselect9(){
  var select = document.getElementById("p24");
  var hiddenDiv = document.getElementById("razon3");
  var input = document.getElementById("xq7")

  if(select.value==="53"){
    input.required=true;
    hiddenDiv.classList.remove("d-none");
  }
  else{
    input.required=false;
    input.value="";
    hiddenDiv.classList.add("d-none");
  }
}

function activaselect10(){
  var select = document.getElementById("p37");
  var hiddenDiv = document.getElementById("razon4");
  var input = document.getElementById("xq8")

  if(select.value==="107"){
    input.required=true;
    hiddenDiv.classList.remove("d-none");
  }
  else{
    input.required=false;
    input.value="";
    hiddenDiv.classList.add("d-none");
  }
}

function activaselect11(){
  var select = document.getElementById("p40");
  var hiddenDiv = document.getElementById("razon5");
  var input = document.getElementById("xq11")

  if(select.value==="66"){
    input.required=false;
    hiddenDiv.classList.remove("d-none");
  }
  else{
    input.required=false;
    input.value="";
    hiddenDiv.classList.add("d-none");
  }
}

function activaselect12(){
  var select = document.getElementById("p41");
  var hiddenDiv = document.getElementById("razon6");
  var input = document.getElementById("xq12")

  if(select.value==="66"){
    input.required=false;
    hiddenDiv.classList.remove("d-none");
  }
  else{
    input.required=false;
    input.value="";
    hiddenDiv.classList.add("d-none");
  }
}

function participo(){
  var select = document.getElementById("p54");
  var hiddenDiv = document.getElementById("participo");
  var input = document.getElementById("p55")

  if(select.value==="66"){
    input.required=true;
    hiddenDiv.classList.remove("d-none");
  }
  else{
    input.required=false;
    input.value="";
    hiddenDiv.classList.add("d-none");
  }
}