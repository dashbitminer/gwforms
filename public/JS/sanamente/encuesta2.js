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

  var temperatureElement = document.querySelector('.temperature');

  function validateNumber(input) {
    if (input.value > 10 || input.value < 0) {
        input.setCustomValidity('El número no puede ser mayor a 10.');
        input.value = "";
        showMessage('Por favor, ingrese un número entre 0 y 10.');
    } else {
        input.setCustomValidity('');
    }
}

function showMessage(message) {
    // Obtener el elemento con el id 'angustia'
    var angustiaElement = document.getElementById('angustia');

    // Crear el HTML para el div
    var divHTML = '<div id="message">' + message + '</div>';

    // Insertar el div inmediatamente después del elemento 'angustia'
    angustiaElement.insertAdjacentHTML('afterend', divHTML);
    
 // Establecer un temporizador para eliminar el div después de 3 segundos
 setTimeout(function() {
  var messageElement = document.getElementById('message');
  if (messageElement) {
      messageElement.remove();
  }
}, 3000);
}

  function updateTemperature(temp) {
    temperatureElement.setAttribute('data-temp', temp);
    var maxHeight = 150; // Altura máxima del termómetro en píxeles
    var height = (temp / 10) * maxHeight; // Ajustar la altura según el valor de temperatura
    temperatureElement.style.height = height + "px";
  }
  function activaselect(){
var select = document.getElementById("participar");
var hiddenDiv = document.getElementById("ocultar"); 
var input1=document.getElementById("trabajo");
var input2=document.getElementById("edad");
var input3=document.getElementById("aniol");
var input4=document.getElementById("sexo");
var input5=document.getElementById("puesto");
var input6=document.getElementById("afirma1");
var input7=document.getElementById("afirma2");
var input8=document.getElementById("afirma3");
var input9=document.getElementById("afirma4");
var input10=document.getElementById("afirma5");
var input11=document.getElementById("afirma6");
var input12=document.getElementById("afirma7");
var input13=document.getElementById("afirma8");
var input14=document.getElementById("afirma9");
var input18=document.getElementById("actitud");
var input19=document.getElementById("estrategia18");
var input20=document.getElementById("apoyo");
var input21=document.getElementById("cambiov");
var input23=document.getElementById("conocet");
var input24=document.getElementById("practicat");
var input25=document.getElementById("estadoe");
var input26=document.getElementById("angustia");
var input27=document.getElementById("oportunidad");
var input28=document.getElementById("clima");
var input29=document.getElementById("conocee");
var input30=document.getElementById("consciente1");
var input31=document.getElementById("beneficio");
var input32=document.getElementById("rel1");
var input33=document.getElementById("rel2");
var input34=document.getElementById("rel3");
var input35=document.getElementById("rel4");
var input37=document.getElementById("afirma1B");
var input38=document.getElementById("afirma2B");
var input39=document.getElementById("afirma3B");
var input40=document.getElementById("afirma4B");
var input41=document.getElementById("afirma5B");
var input42=document.getElementById("afirma6B");
var input44=document.getElementById("afirma7B");
var input44=document.getElementById("afirma8B");
var input45=document.getElementById("afirma9B");
var input46=document.getElementById("rel12");
var input47=document.getElementById("rel22");
var input48=document.getElementById("rel32");
var input49=document.getElementById("rel42");




if(select.value==="1"){ 
hiddenDiv.classList.remove("d-none");
input1.required = true
input2.required = true
input3.required = true
input4.required = true
input5.required = true
input6.required = true
input7.required = true
input8.required = true
input9.required = true
input10.required = true
input11.required = true
input12.required = true
input13.required = true
input14.required = true
input18.required = true
input19.required = true
input23.required = true
input24.required = true
input25.required = true
input26.required = true
input27.required = true
input28.required = true
input29.required = true
input30.required = true
input32.required = true
input33.required = true
input34.required = true
input35.required = true
input37.required = true
input38.required = true
input39.required = true
input40.required = true
input41.required = true
input42.required = true
input44.required = true
input45.required = true
input46.required = true
input47.required = true
input48.required = true
input49.required = true




}
else{
  hiddenDiv.classList.add("d-none");
input1.required = false
input2.required = false
input3.required = false
input4.required = false
input5.required = false
input6.required = false
input7.required = false
input8.required = false
input9.required = false
input10.required = false
input11.required = false
input12.required = false
input13.required = false
input14.required = false
input18.required = false
input19.required = false
input20.required = false
input21.required = false
input23.required = false
input24.required = false
input25.required = false
input26.required = false
input27.required = false
input28.required = false
input29.required = false
input30.required = false
input31.required = false
input32.required = false
input33.required = false
input34.required = false
input35.required = false
input37.required = false
input38.required = false
input39.required = false
input40.required = false
input41.required = false
input42.required = false
input44.required = false
input45.required = false
input46.required = false
input47.required = false
input48.required = false
input49.required = false

}
}

function activaselect2(){
  var select = document.getElementById("consciente1");
  var hiddenDiv = document.getElementById("beneficiodiv"); 
  var input1=document.getElementById("beneficio");
  if(select.checked){ 
    hiddenDiv.classList.remove("d-none");
    input1.required = true
  } 
else{
  hiddenDiv.classList.add("d-none");
  input1.required = false
}
} 

function otro2(){
  var select=document.getElementById('afecta7');
  var linea=document.getElementById("afecta1linea");;
  var select2=document.getElementById('afecta8');
  var linea2=document.getElementById("afecta2linea");;
  var select3=document.getElementById('afecta9');
  var linea3=document.getElementById("afecta3linea");;
  var hiddenDiv4=document.getElementById('afectadiv');
  if(select.checked){
    linea.style.display = "table-row";
  }
else 
{
  linea.style.display = "none";
}

if (select2.checked) {
  linea2.style.display = "table-row";
}
else 
{
  linea2.style.display = "none";
}

if (select3.checked){
  linea3.style.display = "table-row";
}

else 

{
  linea3.style.display = "none";
}

if(select.checked||select2.checked||select3.checked){
  hiddenDiv4.classList.remove("d-none");
}

else{
  hiddenDiv4.classList.add("d-none");
}
}

function otro(){
  var select=document.getElementById('estrategia18');
  var select2=document.getElementById('estrategia19');
  var select3=document.getElementById('estrategia20');
  var hiddenDiv4=document.getElementById('epracticadiv');
  if (select.checked){
    hiddenDiv4.classList.remove("d-none");
  }
  
  if (select2.checked){
    hiddenDiv4.classList.remove("d-none");
  }
  
  if (select3.checked){
    hiddenDiv4.classList.add("d-none");
  }
}

function otro4(){
  var select3=document.getElementById('ap201');
  var hiddenDiv4=document.getElementById('epracticadiv');
  var hiddenDiv5=document.getElementById('motivodiv');
  if (select3.checked){
    hiddenDiv5.classList.remove("d-none");
  }
  else{
    hiddenDiv5.classList.add("d-none");
  }

  if (select2.checked){
    hiddenDiv4.classList.remove("d-none");
  }  
}

function otro5(){
  var select3=document.getElementById('opractica39');
  var select4=document.getElementById('opractica38');
  var hiddenDiv5=document.getElementById('8p');
  var hiddenDiv4=document.getElementById('82div');
  if (select3.checked){
    hiddenDiv5.classList.add("d-none");
  }
  else{
    hiddenDiv5.classList.remove("d-none");
  }
  if (select4.checked){
    hiddenDiv4.classList.remove("d-none");
  }
  else{
    hiddenDiv4.classList.add("d-none");
  }
}

function otro6(){
  var select=document.getElementById('practicat1');
  var hiddenDiv4=document.getElementById('831div');
  if (select.checked){
    hiddenDiv4.classList.remove("d-none");
  }
  else{
    hiddenDiv4.classList.add("d-none");
  }

}
function otro3(){
  var select=document.getElementById('epractica26');
  var hiddenDiv4=document.getElementById('apoyodiv');

  if (select.checked){
    hiddenDiv4.classList.remove("d-none");
 }
  else{
    hiddenDiv4.classList.add("d-none");
  }
}

document.addEventListener("DOMContentLoaded", function() {
  const images = document.querySelectorAll(".img-max");

  images.forEach(function(image) {
    image.addEventListener("click", function() {
      const radio = image.closest(".form-check").querySelector(".form-check-input");
      radio.checked = true;
    });
  });
});