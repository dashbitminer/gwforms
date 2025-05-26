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

function activaselect2(){
var select = document.getElementById("Eactual");
var hiddenDiv = document.getElementById("hidden-question-div");
var hiddenDiv2 = document.getElementById("hidden-question-div2");
var hiddenDiv3 = document.getElementById("hidden-question-div3");
var hiddenDiv4 = document.getElementById("hidden-question-div5");
var input = document.getElementById("UT");
var input2 = document.getElementById("GA");
var input3 = document.getElementById("Turno");
var input4 = document.getElementById("seccion");
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
    hiddenDiv3.classList.add("d-none");
  input3.required = false;
      hiddenDiv4.classList.add("d-none");
  input4.required = false;
  }

  }

function activaselect3(){
var hiddenDiv3 = document.getElementById("hidden-question-div3");
var hiddenDiv4 = document.getElementById("hidden-question-div5");
var input2 = document.getElementById("GA");
var input3 = document.getElementById("Turno");
var input4 = document.getElementById("seccion");
  if(input2.value==="38"||input2.value==="39"||input2.value==="40"){ 
  hiddenDiv3.classList.add("d-none");
  input3.required = false;
      hiddenDiv4.classList.add("d-none");
  input4.required = false;}
  else{
  hiddenDiv3.classList.remove("d-none");
  input3.required = false;
    hiddenDiv4.classList.remove("d-none");
  input4.required = false;
  }
  }

  function ver1(){
    var div= document.getElementById("autoRes");
    div.classList.remove("d-none");
  }
  
  function ocultar1(){
    var div= document.getElementById("autoRes");
    div.classList.add("d-none");
  }
  
  function activaselect4(){
  var select = document.getElementById("FechaN");
  var hiddenDiv = document.getElementById("participaAut");
  var hiddenDiv1 = document.getElementById("infoNNA");
  var birthDate = new Date(select.value);
  var diff_ms = Date.now() -birthDate ;
  var age_dt = new Date(diff_ms); 
  var age= Math.abs(age_dt.getUTCFullYear() - 1970);
      hiddenDiv.classList.remove("d-none");
  var input = document.getElementById("autR10");
  var input1= document.getElementById("autR11");
  
    hiddenDiv.classList.add("d-none");
    hiddenDiv1.classList.add("d-none");

    if(age>=12){    
    hiddenDiv.classList.remove("d-none");
    input.required = true;
    input1.required = true;
    } else {
    hiddenDiv1.classList.remove("d-none");
    input.required = false;
    input1.required = false;
  
    }
  }

  function ver3(){
    var div= document.getElementById("infoNNA");
    var div1= document.getElementById("autPart");
    div.classList.remove("d-none");
    div1.classList.remove("d-none");
  }
  
  function ocultar3(){
    var div= document.getElementById("infoNNA");
    var div1= document.getElementById("autPart");
    div.classList.add("d-none");
    div1.classList.add("d-none");
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

function activaselect51(){
  var select = document.getElementById("nacionalidad");
  var labelOutput = document.getElementById("DUIL");
  if(select.value==="1"){    
    labelOutput.innerHTML = "Número de CURP:*"
    document.getElementById("DUIM").placeholder = "abcd000000abcdef00";
  } else {labelOutput.innerHTML = "Número de tu Documento de Identidad"
    document.getElementById("DUIM").placeholder = ""
  } 
} 

  function validateFormat() {
    var formatInput = document.getElementById("DUIM").value;
    var output2 = document.getElementById("duil");
      output2.innerHTML = formatInput;
  }

  function actualizarCiudades(paisSeleccionado) {
    var ciudadSelect = document.getElementById("departamento");
    ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades

    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesCiudades = this.responseText;
            ciudadSelect.innerHTML = opcionesCiudades; // Actualizar la lista de ciudades
        }
    };
    xhr.open("POST", "/obtenerciudades", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("pais=" + paisSeleccionado);
}

function actualizarmun(paisSeleccionado) {
  var ciudadSelect = document.getElementById("departamento2");
  ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
  var sede = document.getElementById("sede2");
  sede.innerHTML = ""; // Limpiar la lista de ciudades
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var pcionesmunsede= this.responseText;
          ciudadSelect.innerHTML = pcionesmunsede; // Actualizar la lista de ciudades
      }
  };
  xhr.open("POST", "/obtenermunsedeedu", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("depsede=" + paisSeleccionado);
}

function actualizarsede(paisSeleccionado) {
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
  xhr.open("POST", "/obtenersedeedu", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("munsede=" + paisSeleccionado);
}

function guardar2(event){
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
              alert("Se guardó correctamente");
              // redireccionar a la página de inicio
              window.location.href = "/educacion/inscripcion/nna/mx";
          }else if (respuesta=="-1"){
            alert("Su respuesta ya se encuentra registrada");
              // redireccionar a la página de inicio
              window.location.href = "/educacion/inscripcion/nna/mx";
          }           
          else {
              alert("Hubo un error al guardar, intente de nuevo");
          }
      }
  };
  xhr.open("POST", "/educacion/inscripcion/nna/post", true);
  xhr.send(formData);
  }

  function guardar1(event){
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
                alert("Se guardó correctamente");
                // redireccionar a la página de inicio
                window.location.href = "/educacion/inscripcion/nna/rd";
            }else if (respuesta=="-1"){
              alert("Su respuesta ya se encuentra registrada");
                // redireccionar a la página de inicio
                window.location.href = "/educacion/inscripcion/nna/rd";
            }           
            else {
                alert("Hubo un error al guardar, intente de nuevo");
            }
        }
    };
    xhr.open("POST", "/educacion/inscripcion/nna/post", true);
    xhr.send(formData);
    }