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

function hideAlert2() {
  var alertContainer = document.getElementById("alertContainer2");
  alertContainer.style.display = "none";
}

function activaselect2(){
  var select = document.getElementById("Eactual");
var hiddenDiv = document.getElementById("hidden-question-div");
var hiddenDiv2 = document.getElementById("hidden-question-div2");
var hiddenDiv4 = document.getElementById("hidden-question-div5");
var input = document.getElementById("UT");
var input2 = document.getElementById("GA");
var input3 = document.getElementById("Turno");
  if(select.value==="1"){    
      hiddenDiv.classList.add("d-none");
      input.required = false;
      hiddenDiv2.classList.remove("d-none");
      input2.required = true;
      
      activaselect3();
  }
  
   else {
  hiddenDiv.classList.remove("d-none");
  input.required = true;
    hiddenDiv2.classList.add("d-none");
  input2.required = false;
  input3.required = false;
      hiddenDiv4.classList.add("d-none");
  }
  }

function activaselect3(){
var hiddenDiv4 = document.getElementById("hidden-question-div5");
var input2 = document.getElementById("GA");
var input3 = document.getElementById("Turno");
var input4 = document.getElementById("seccion");
  if(input2.value==="38"||input2.value==="39"||input2.value==="40"){ 
  input3.required = false;
      hiddenDiv4.classList.add("d-none");
  input4.required = false;}
  else{
  input3.required = true;
    hiddenDiv4.classList.remove("d-none");
  input4.required = true;
  }
  }

function activaselect4(){
var select = document.getElementById("FechaN");
var hiddenDiv = document.getElementById("llenadofecha");
var hiddenDiv2 = document.getElementById("menoredad");
var hiddenDiv3 = document.getElementById("mayorpart");
var hiddenDiv4 = document.getElementById("mayorpart2");
var hiddenDiv5 = document.getElementById("duiparticipante");
var output1 = document.getElementById("edadl");
var birthDate = new Date(select.value);
var diff_ms = Date.now() -birthDate ;
var age_dt = new Date(diff_ms); 
var age= Math.abs(age_dt.getUTCFullYear() - 1970);
var input = document.getElementById("aut4");
var input2 = document.getElementById("aut3");
var input3 = document.getElementById("aut2");
var input4 = document.getElementById("DUIM");
var input5 = document.getElementById("consentimiento");
     
  if(age >= 18 && age <= 35){    
    hiddenDiv2.classList.add("d-none");
    hiddenDiv3.classList.remove("d-none");
    hiddenDiv4.classList.remove("d-none");
    hiddenDiv.classList.remove("d-none");
    output1.innerHTML = age;
    input.required = true;
    input2.required = true;
    input3.required = true;
    input4.required = true;
    input5.required = false;
  } else if(age < 16 || age > 35) {
    hiddenDiv.classList.add("d-none");
    hiddenDiv2.classList.add("d-none");
    hiddenDiv3.classList.add("d-none");
    var alertContainer = document.getElementById("alertContainer2");
    alertContainer.innerHTML = '<div class="alert alert-danger" role="alert">Pareces que no estas dentro del rango de edad para participar en el programa que es de 16 a 29 años. Si deseas recibir mayor información, puedes contactar con el líder del programa.</div>';
    alertContainer.style.display = "block";
  } else {
    hiddenDiv.classList.add("d-none");
    hiddenDiv2.classList.remove("d-none");
    hiddenDiv3.classList.add("d-none");
    hiddenDiv4.classList.add("d-none");
    hiddenDiv5.classList.add("d-none");
    input.required = false;
    input2.required = false;
    input3.required = false;
    input4.required = false;
    input5.required = true;
    var nacionalidad = document.getElementById('nacionalidad');
    nacionalidad.value = '';
  }
  }

  function activaselect5(){
    var select = document.getElementById("sede");
    var hiddenDiv = document.getElementById("diplomados");
    
    if(select.value==="666"){
      hiddenDiv.classList.remove("d-none") ;
    }
    else if(select.value==="2968"){
      hiddenDiv.classList.remove("d-none") ;
    }
    else {
       hiddenDiv.classList.add("d-none") ;
    }
    }

    
    function activaselect8(){
    var select = document.getElementById("WS");
    var hiddenDiv = document.getElementById("what");
     input4 = document.getElementById("TM2")
    if(select.value==="1"){
      hiddenDiv.classList.remove("d-none") ;
      input4.required = true;
    }
    else {
       hiddenDiv.classList.add("d-none") ;
       input4.required = false;
    }
    }
   
function activaselect10(){
  var select = document.getElementById("dis");
var hiddenDiv = document.getElementById("hidden-question-div10");
  if(select.value==="1"){    
  hiddenDiv.classList.remove("d-none");
  } else {
  hiddenDiv.classList.add("d-none");
    }
  }


  function validateFormat() {
    var formatInput = document.getElementById("DUIM").value;
       var formatRegex = /^\d{8}-\d$/;
       var select = document.getElementById("nacionalidad");
      if(select.value==="1"){
            if (!formatRegex.test(formatInput)) {
          var alertContainer = document.getElementById("alertContainer");
          alertContainer.innerHTML = '<div class="alert alert-danger" role="alert">El formato del DUI no es válido.</div>';
          alertContainer.style.display = "block";
          return false;
    }
    return true;
  }
  }

      function hideAlert() {
    var alertContainer = document.getElementById("alertContainer");
    alertContainer.style.display = "none";
  }

  document.querySelector('#formualrioRegistro').addEventListener('submit', function(event) {
  var interests = document.querySelectorAll('input[name="dis2"]:checked');
   var select = document.getElementById("dis");
  if(select.value==="1"){
  if (interests.length === 0) {
    alert('Debes seleccionar al menos una opción de discapacidad en caso de poseer alguna');
    event.preventDefault();
  }
  }
});

function activa1(){
  var input = document.getElementById("nombre");
  var input2 = document.getElementById("apellido");
  var output = document.getElementById("nombrep");
    output.innerHTML = input.value +" "+input2.value;
  }

  function activaselect51(nacionalidad){
    var select = document.getElementById("FechaN");
    var birthDate = new Date(select.value);
  var diff_ms = Date.now() -birthDate ;
  var age_dt = new Date(diff_ms); 
  var age= Math.abs(age_dt.getUTCFullYear() - 1970);
  
    var input = document.getElementById("extranjero");
    var input2 = document.getElementById("duiparticipante");
  
     if (nacionalidad == "1" && age >= 18 && age <= 29){
      // añadir placeholder de variable id DUIM
      document.getElementById("DUIM").placeholder = "00000000-0";
      document.getElementById("divextranjero").classList.add("d-none");
      document.getElementById("duiparticipante").classList.remove("d-none");
      input.required = false;
      input2.required = true;
    } else if (age >= 18 && age <= 29 && nacionalidad == "0"){
      document.getElementById("DUIM").placeholder = "";
      document.getElementById("divextranjero").classList.remove("d-none");
      document.getElementById("duiparticipante").classList.remove("d-none");
      input.required = true;
      input2.required = true;
    } else if((age >= 16 || age <18) && nacionalidad == "0") {
      document.getElementById("duiparticipante").classList.add("d-none");
      document.getElementById("divextranjero").classList.remove("d-none");
      input.required = false;
      input2.required = true;
    } else {
      document.getElementById("duiparticipante").classList.add("d-none");
      document.getElementById("divextranjero").classList.add("d-none");
      input.required = false;
      input2.required = false;
    }
  
  }

function verificar(consentimiento){
  if(consentimiento=== "1"){
    document.getElementById("negativoconsentimiento").classList.add("d-none");
    document.getElementById("llenadofecha").classList.remove("d-none");
  } else {
    document.getElementById("negativoconsentimiento").classList.remove("d-none");
    document.getElementById("llenadofecha").classList.add("d-none");
  }
}

function transporte(transporte){
  if(transporte=== "4"){
    document.getElementById("hidden-question-div4").classList.add("d-none");
    document.getElementById("gasto").required = false;
  } else {
    document.getElementById("hidden-question-div4").classList.remove("d-none");
    document.getElementById("gasto").required = true;
  }
}

function sedeopciones(sede){
  if(sede=== "2968"){
    document.getElementById("mentoriasdiv").classList.add("d-none");
    document.getElementById("dispositivodiv").classList.add("d-none");
    document.getElementById("dispositivo").required = false;
    document.getElementById("invalidCheck").required = false;
  } else {
    document.getElementById("mentoriasdiv").classList.remove("d-none");
    document.getElementById("dispositivodiv").classList.remove("d-none");
    document.getElementById("dispositivo").required = true;
    document.getElementById("invalidCheck").required = true;
  }
}