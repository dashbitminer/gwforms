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

        if (!validarClubs()) {
          event.preventDefault()
          event.stopPropagation()
      }
  
        form.classList.add('was-validated')
      }, false)
    })
  })()

  var signaturePad;
  var signaturePad2;
  
  function signiture(){
    var input = document.getElementById("Signature1draw");
    signaturePad = new SignaturePad(input);
    var input2 = document.getElementById("Signature2draw");
    signaturePad2 = new SignaturePad(input2);
  }
  
  function clearSigniture(){
    signaturePad.clear();
  }
  
  function clearSigniture2(){
    signaturePad2.clear();
  }
  
  document.addEventListener('DOMContentLoaded',signiture);

  function activa1(){
    var input = document.getElementById("Firstmiddle");
    var input2 = document.getElementById("Surnames");
    var input3 = document.getElementById("sede2");

    var output = document.getElementById("nombrep");
    var output2 = document.getElementById("nombrep2");
    var output4 = document.getElementById("nombreSchool1")

    var output5 = document.getElementById("nombrePart2");
    var output7 = document.getElementById("nombreSchool2")

      output.innerHTML = input.value +" "+input2.value;
      output2.innerHTML = input.value +" "+input2.value;
      output4.innerHTML = input3.options[input3.selectedIndex].text //Para mostrar el texto en lugar del id/valor del select
      output5.innerHTML = input.value +" "+input2.value;
      output7.innerHTML = input3.options[input3.selectedIndex].text //Para mostrar el texto en lugar del id/valor del select
    }
  
function activaselect4(){
var select = document.getElementById("FechaN");
var hiddenDiv1 = document.getElementById("menor1");
var hiddenDiv3 = document.getElementById("participante");

var input1 = document.getElementById("autR1");
var input2 = document.getElementById("autR2");
var input3 = document.getElementById("autR3");
var input4 = document.getElementById("nombreRes");
var input5 = document.getElementById("Telcontacto");
var input6 = document.getElementById("relacion");
var input7 = document.getElementById("autR11");

var birthDate = new Date(select.value);
var diff_ms = Date.now() -birthDate ;
var age_dt = new Date(diff_ms); 
var age= Math.abs(age_dt.getUTCFullYear() - 1970);
 
  if(age>=18){    
  hiddenDiv1.classList.add("d-none");
  hiddenDiv3.classList.remove("d-none");

    input1.required = false;
    input2.required = false;
    input3.required = false;
    input4.required = false;
    input5.required = false;
    input6.required = false;
    input7.required = false;

} else {
    hiddenDiv1.classList.remove("d-none");
    hiddenDiv3.classList.add("d-none");
    
    input1.required = true;
    input2.required = true;
    input3.required = true;
    input4.required = true;
    input5.required = true;
    input6.required = true;
    input7.required = true;
    
    }
  }

  function ver1(){
    var div= document.getElementById("autorizacionesr2");
    var div2= document.getElementById("participante");
    var div3= document.getElementById("comentarior1");

    div.classList.remove("d-none");
    div2.classList.remove("d-none");
    div3.classList.add("d-none");
  }
  
  function ocultar1(){
    var div= document.getElementById("autorizacionesr2");
    var div2= document.getElementById("participante");
    var div3= document.getElementById("comentarior1");

    div.classList.add("d-none");
    div2.classList.add("d-none");
    div3.classList.remove("d-none");
  }

  function ver2(){
    var div= document.getElementById("autorizacionesr3");
    var div2= document.getElementById("participante");
    var div3= document.getElementById("comentarior2");

    div.classList.remove("d-none");
    div2.classList.remove("d-none");
    div3.classList.add("d-none");
  }
  
  function ocultar2(){
    var div= document.getElementById("autorizacionesr3");
    var div2= document.getElementById("participante");
    var div3= document.getElementById("comentarior2");

    div.classList.add("d-none");
    div2.classList.add("d-none");
    div3.classList.remove("d-none");
  }

  function ver3(){
    var div= document.getElementById("mediosr1");
    div.classList.remove("d-none");

    var input= document.getElementById("FechaSt1");
    var input2= document.getElementById("FechaEnd1");
    
    input.required = true;
    input2.required = true;    
  }
  
  function ocultar3(){
    var div= document.getElementById("mediosr1");
    div.classList.add("d-none");

    var input= document.getElementById("FechaSt1");
    var input2= document.getElementById("FechaEnd1");
    
    input.required = false;
    input2.required = false; 
  }

  function ver4(){
    var div= document.getElementById("autorizacionesp1");
    var div2= document.getElementById("comentarior4");

    div.classList.remove("d-none");
    div2.classList.add("d-none")
  }
  
  function ocultar4(){
    var div= document.getElementById("autorizacionesp1");
    var div2= document.getElementById("comentarior4");

    div.classList.add("d-none");
    div2.classList.remove("d-none")
  }

  function ver5(){
    var div= document.getElementById("autorizacionesp2");
    var div2= document.getElementById("comentarior5");

    div.classList.remove("d-none");
    div2.classList.add("d-none")
  }
  
  function ocultar5(){
    var div= document.getElementById("autorizacionesp2");
    var div2= document.getElementById("comentarior5");

    div.classList.add("d-none");
    div2.classList.remove("d-none")
  }

  function ver6(){
    var div= document.getElementById("mediosp1");
    div.classList.remove("d-none");
  
    var input= document.getElementById("FechaSt2");
    var input2= document.getElementById("FechaEnd2");
    
    input.required = true;
    input2.required = true;    
}
  
  function ocultar6(){
    var div= document.getElementById("mediosp1");
    div.classList.add("d-none");

    var input= document.getElementById("FechaSt2");
    var input2= document.getElementById("FechaEnd2");
    
    input.required = false;
    input2.required = false;   
}

function validarClubs() {
  var checkboxes = document.querySelectorAll('input[name="club"]');
  var alMenosUnoSeleccionado = false;
  var club = document.getElementById("nombreClub2");
  club.innerHTML = "";

  checkboxes.forEach(function(checkbox) {
    switch (checkbox.value) {
      case "1":
       clubtext = "Reading club";
       break;
      case "2":
        clubtext = "Peer mentoring";
        break;
      case "3":
        clubtext = "Relaxation or practice club";
        break;
    }
      if (checkbox.checked) {
          alMenosUnoSeleccionado = true;
      if (club.innerHTML == ""){
        club.innerHTML = clubtext;
      }
      else{
        club.innerHTML = club.innerHTML + ", " + clubtext;
      }
      }
  });

  if (!alMenosUnoSeleccionado) {
      alert('You must select at least one club.');
      return false; // Esto evita que el formulario se envíe si ningún checkbox está seleccionado.
  }
  return true; // Si al menos uno está seleccionado, permite el envío del formulario.
}

function guardar(event){
  event.preventDefault(); // Evitar el envío del formulario
  var form = document.getElementById("formualrioRegistro");
  var formData = new FormData(form);
  // verificar los campos requeridos
  if (form.checkValidity() === false) {
      form.reportValidity();
      event.stopPropagation();
      return;
  }
  if (typeof signaturePad !== "undefined" && signaturePad && !signaturePad.isEmpty()) {
      var firma1 = signaturePad.toDataURL();
      formData.append("Signature1draw", firma1);
  } 
  
  if (typeof signaturePad2 !== "undefined" && signaturePad2 && !signaturePad2.isEmpty()) {
      var firma2 = signaturePad2.toDataURL();
      formData.append("Signature2draw", firma2);
  } 
  // Enviar petición AJAX
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var respuesta = this.responseText;
          if(respuesta=="1"){
              alert("Saved successfully");
              // redireccionar a la página de inicio
              window.location.href = "/educacion/ny/eng";
          }else if (respuesta=="-1"){
            alert("Your response is already registered");
              // redireccionar a la página de inicio
              window.location.href = "/educacion/ny/eng";
          }           
          else {
              alert("There was an error saving, please try again");
          }
      }
  };
  xhr.open("POST", "/educacion/ny/post", true);
  xhr.send(formData);
  }