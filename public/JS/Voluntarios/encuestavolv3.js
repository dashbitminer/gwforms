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

function activaselectacti(checkbox) {
 const inputEdu = document.getElementById('p62edu');
const inputSalud = document.getElementById('p62salud');
const inputSA = document.getElementById('p62sa');

const inputs = [inputEdu, inputSalud, inputSA];

// Mostrar/ocultar cada div individualmente según su checkbox
const c14 = document.querySelector('input[name="p3"][value="14"]');
const c15 = document.querySelector('input[name="p3"][value="15"]');
const c130 = document.querySelector('input[name="p3"][value="130"]');

const d14 = document.getElementById('actividadesocultar1');
const d15 = document.getElementById('actividadesocultar2');
const d130 = document.getElementById('actividadesocultar3');

// Mostrar u ocultar divs según checkbox
d14.classList.toggle("d-none", !c14.checked);
d15.classList.toggle("d-none", !c15.checked);
d130.classList.toggle("d-none", !c130.checked);

// Requerir el select correspondiente si su checkbox está activo
if (c14.checked) inputEdu.required = true; else inputEdu.required = false;
if (c15.checked) inputSalud.required = true; else inputSalud.required = false;
if (c130.checked) inputSA.required = true; else inputSA.required = false;

// Si ninguno está seleccionado, limpia los valores
const anyChecked = [c14, c15, c130].some(cb => cb.checked);
if (!anyChecked) {
  inputs.forEach(input => {
    if (input) input.value = "";
  });
}
}


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

function activaselect3(select){ //actual por cambio en form
var hiddenDiv = document.getElementById("otrol3");
  var input = document.getElementById("otro3");
  if (!select || !hiddenDiv || !input) return;

  if(select.value === "42"){
    input.required = true;
    hiddenDiv.classList.remove("d-none");
  } else {
    input.required = false;
    input.value = "";
    hiddenDiv.classList.add("d-none");
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
  var input = document.getElementById("apoyo")

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

function activaselect6(checkbox) {
  var hiddenDiv1 = document.getElementById('otrol5');
  var input = document.getElementById('apoyo');
  
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
  var select = document.getElementById("p56");
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

function activaselect13(){
  var select = document.getElementById("p38");
  var hiddenDiv = document.getElementById("razon7");
  var input = document.getElementById("xq9")

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

function activaselect14(){
  var select = document.getElementById("p39");
  var hiddenDiv = document.getElementById("razon8");
  var input = document.getElementById("xq10")

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

function activaselect15(){
  var select = document.getElementById("p59");
  var hiddenDiv = document.getElementById("razon9");
  var input = document.getElementById("xq1")

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


function activaselectcheckboxes() {
  const seleccion = document.getElementById("p2").value;
  const checkboxes = document.querySelectorAll('[data-categorias]');

  checkboxes.forEach(cb => {
    const categorias = cb.getAttribute("data-categorias").split(" ");
    if (categorias.includes(seleccion)) {
      cb.style.display = "block";
    } else {
      cb.style.display = "none";
      cb.querySelector("input").checked = false;
    }
  });
}

// Ejecutar al cargar la página si quieres ocultarlos inicialmente
window.addEventListener("DOMContentLoaded", activaselectcheckboxes);

