// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
  'use strict'

  // Fetch all the forms we want to apply custom Bootstrap validation styles to
  const forms = document.querySelectorAll('.needs-validation')

  // Loop over them and prevent submission
  Array.from(forms).forEach(form => {
    form.addEventListener('submit', event => {
      validateFormat()
      if(verificarSeleccion() === false){
        event.preventDefault();
      }
      if (!form.checkValidity()) {
        event.preventDefault()
        event.stopPropagation()
      }
       else if (document.getElementById("alertContainer").style.display === "block") {
    event.preventDefault();
    alert("No se puede enviar el formulario mientras este incorrecto el DNI");
  }
    }, false)
  })
})()


  function hideAlert() {
    var alertContainer = document.getElementById("alertContainer");
    alertContainer.style.display = "none";
  }

  function hideAlert2() {
    var alertContainer = document.getElementById("alertContainer2");
    alertContainer.style.display = "none";
  }
 
function activaselect2(){
  var select = document.getElementById("TdC");
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
  
function activaselect4(){
var select = document.getElementById("FechaN");
var hiddenDiv = document.getElementById("ocultar");
var hiddenDiv2 = document.getElementById("menor");
var hiddenDiv3 = document.getElementById("hidden-question-div6");
var hiddenDiv4 = document.getElementById("brindo2");
var hiddenDiv5 = document.getElementById("mayorpart");
var hiddenDiv6 = document.getElementById("mayorpart2");
var hiddenDiv7 = document.getElementById("duiparticipante");
var hiddenDiv8 = document.getElementById("brindo");
var birthDate = new Date(select.value);
var diff_ms = Date.now() -birthDate ;
var age_dt = new Date(diff_ms); 
var age= Math.abs(age_dt.getUTCFullYear() - 1970);
     hiddenDiv.classList.remove("d-none");
var input = document.getElementById("DUIM");
var output1 = document.getElementById("edadl");
var input9= document.getElementById("nombreCompleto");
var input10= document.getElementById("apellido");
var input11= document.getElementById("Sexo");
var input12= document.getElementById("pais");
var input13= document.getElementById("departamento");
var input22= document.getElementById("tipo");
var input23= document.getElementById("TdC");
var input28= document.getElementById("dis");
var input30= document.getElementById("gp");
var input41= document.getElementById("aut4");
var input42= document.getElementById("aut2");
var input43= document.getElementById("aut3");
var input50= document.getElementById("nacionalidad");
 
  if(age>=18){    
  hiddenDiv2.classList.add("d-none");
  hiddenDiv3.classList.remove("d-none");
  hiddenDiv4.classList.remove("d-none");
  hiddenDiv5.classList.remove("d-none");
  hiddenDiv6.classList.remove("d-none");
  hiddenDiv7.classList.remove("d-none");
  hiddenDiv8.classList.add("d-none");
  output1.innerHTML = age;
  input.required = true;

input9.required = true;
input10.required = true;
input11.required = true;
input12.required = true;
input13.required = true;
input22.required = true;
input23.required = true;
input28.required = true;
input30.required = true;
input41.required = true;
input42.required = true;
input43.required = true;
input50.required = true;
 }
 else{
  var alertContainer = document.getElementById("alertContainer2");
  alertContainer.innerHTML = '<div class="alert alert-danger" role="alert">Parece que has escrito mal tu fecha de nacimiento, por favor revisar.</div>';
  alertContainer.style.display = "block";

  hiddenDiv.classList.add("d-none") ;
  hiddenDiv2.classList.add("d-none") ;
  hiddenDiv3.classList.remove("d-none") ;
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

function activaselect51(){
var select = document.getElementById("nacionalidad");
var labelOutput = document.getElementById("DUIL");
if(select.value==="1"){    
  labelOutput.innerHTML = "Escribe el número documento personal de identidad (DNI)*"
  document.getElementById("DUIM").placeholder = "0000-0000-00000";
    } else {labelOutput.innerHTML = "Escribe el número de tu Documento de Identidad"
    document.getElementById("DUIM").placeholder = ""
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

function activaselect52(){
var select = document.getElementById("aut");
var hiddenDiv = document.getElementById("brindo");
var hiddenDiv2 = document.getElementById("brindo2");
var hiddenDiv3 = document.getElementById("brindo3");
var input9= document.getElementById("nombreCompleto");
var input10= document.getElementById("apellido");
var input11= document.getElementById("Sexo");
var input12= document.getElementById("pais");
var input13= document.getElementById("departamento");
var input22= document.getElementById("tipo");
var input23= document.getElementById("TdC");
var input28= document.getElementById("dis");
var input30= document.getElementById("gp");
var input45= document.getElementById("docF");
var input46= document.getElementById("docA");
var input50= document.getElementById("nacionalidad");


if(select.value==="1"){
  hiddenDiv.classList.remove("d-none") ;
  hiddenDiv2.classList.remove("d-none") ;
  hiddenDiv3.classList.add("d-none") 
 
  input9.required = true;
input10.required = true;
input11.required = true;
input12.required = true;
input13.required = true;
input22.required = true;
input23.required = true;
input28.required = true;
input30.required = true;
input45.required = true;
input46.required = true;
input50.required = true;

}
else {
   hiddenDiv.classList.add("d-none") ;
   hiddenDiv2.classList.add("d-none") ;
   hiddenDiv3.classList.remove("d-none") ;
   
   input9.required = false;
   input10.required = false;
   input11.required = false;
   input12.required = false;
   input13.required = false;
   input22.required = false;
   input23.required = false;
   input28.required = false;
   input30.required = false;
   input45.required = false;
  input46.required = false;  
  input50.required = false;
}
}

function activa1(){
  var input = document.getElementById("nombreCompleto");
  var input2 = document.getElementById("apellido");
  var output2 = document.getElementById("nombrep2");
    output2.innerHTML = input.value +" "+input2.value;
  }
  
  function validateFormat() {
    var formatInput = document.getElementById("DUIM").value;
    var formatRegex = /^\d{4}-\d{4}-\d{5}$/;
    var select = document.getElementById("nacionalidad");

   if(select.value==="1"||select.value===""){
         if (!formatRegex.test(formatInput)) {
       var alertContainer = document.getElementById("alertContainer");
       alertContainer.innerHTML = '<div class="alert alert-danger" role="alert">El formato del DNI no es válido.</div>';
       alertContainer.style.display = "block";
       return false;
 }
 return true;
    }
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
  $('#sede2').selectize()[0].selectize.destroy(); 
  var sede = document.getElementById("sede2");
  sede.innerHTML = "";
  // Limpiar la lista de ciudades
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var pcionesmunsede= this.responseText;
          ciudadSelect.innerHTML = pcionesmunsede; // Actualizar la lista de ciudades
      }
  };
  xhr.open("POST", "/obtenermunsedesanamente", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("depsede=" + paisSeleccionado);
}

function actualizarsede(paisSeleccionado) {
  var ciudadSelect = document.getElementById("sede2");
  ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
  $('#sede2').selectize()[0].selectize.destroy(); 
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var opcionesmunsede= this.responseText;
          ciudadSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
          $('#sede2').selectize({
            sortField: 'text'
        });
        }
  };
  xhr.open("POST", "/obtenersedesanamente", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("munsede=" + paisSeleccionado);
}

function activaselect5(){
  var select = document.getElementById("tipo");
  var hiddenDiv = document.getElementById("acc2d");
  var input4 = document.getElementById("acc2");
  if(select.value==="10"){    
    hiddenDiv.classList.remove("d-none");
    input4.required = true;
    } else {
    hiddenDiv.classList.add("d-none");
    input4.required = false;
      }
    }

    function activaselect6(){
      var select = document.getElementById("acc2");
      var select2 = document.getElementById("idp");
      var hiddenDiv = document.getElementById("acc3d");
      var input4 = document.getElementById("acc3");
      if(select.value==="1"||select.value==="2"||select.value==="3"){    
        hiddenDiv.classList.remove("d-none");
        input4.required = true;
        input4.innerHTML = ""; // Limpiar la lista de ciudades
    
        // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
        var xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                var opcionesCiudades = this.responseText;
                input4.innerHTML = opcionesCiudades; // Actualizar la lista de ciudades
            }
        };
        xhr.open("POST", "/obtenerperfil", true);
        xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhr.send("pais=" + select.value +"&idp="+select2.value);
      } else if(select.value==="14"||select.value==="15"){
        hiddenDiv.classList.add("d-none");
        input4.required = false;
        var hiddenDiv2 = document.getElementById("acc4d");
        var input5 = document.getElementById("acc4");
        hiddenDiv2.classList.remove("d-none");
        input5.required = true;
    
        var xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                var opcionesCiudades = this.responseText;
                input5.innerHTML = opcionesCiudades; // Actualizar la lista de ciudades
            }
        };
        xhr.open("POST", "/obtenersubperfil", true);
        xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhr.send("pais=" + select.value);
    
     }  else {
        hiddenDiv.classList.add("d-none");
        input4.required = false;
          } 
        }
        
    function activaselect7(){
      var select = document.getElementById("acc3");
      var hiddenDiv = document.getElementById("acc4d");
      var input4 = document.getElementById("acc4");
      var labelOutput = document.getElementById("acc4l");
      if(select.value==="1"){    
        labelOutput.innerHTML = "Selecciona tu rango/categoría:"
        hiddenDiv.classList.remove("d-none");
        input4.required = true;
        hiddenDiv.classList.remove("d-none");
        input4.required = true;
        input4.innerHTML = ""; // Limpiar la lista de ciudades
    
        // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
        var xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                var opcionesCiudades = this.responseText;
                input4.innerHTML = opcionesCiudades; // Actualizar la lista de ciudades
            }
        };
        xhr.open("POST", "/obtenersubperfil", true);
        xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhr.send("pais=" + select.value);
    
        } else if(select.value==="6"||select.value==="10"){
        labelOutput.innerHTML = "Selecciona tu perfil de personal de salud"
         hiddenDiv.classList.remove("d-none");
        input4.required = true;
        hiddenDiv.classList.remove("d-none");
        input4.required = true;
        input4.innerHTML = ""; // Limpiar la lista de ciudades
    
        // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
        var xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                var opcionesCiudades = this.responseText;
                input4.innerHTML = opcionesCiudades; // Actualizar la lista de ciudades
            }
        };
        xhr.open("POST", "/obtenersubperfil", true);
        xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhr.send("pais=" + "2");
        } 
        else {
        hiddenDiv.classList.add("d-none");
        input4.required = false;
          }
        }

        function verificarSeleccion() {
          var select = document.getElementById('sede2');
          if (select.value === '') {
              alert('Por favor, selecciona una sede.');
              return false;
          }
          return true;
      }

      function mostrardivDGDP(){
        var select = document.getElementById("acc2");
        var hiddenDiv = document.getElementById("infoDGDP");
        var input1 =document.getElementById("sedehn");
        var input2 =document.getElementById("tiposedehn");
        var input3 =document.getElementById("cargosedehn");
        var input4 =document.getElementById("depto");
        var input5 =document.getElementById("munic");
        var input6 =document.getElementById("aldeahn");
        var input7 =document.getElementById("caseriohn");
        var input8 =document.getElementById("codigohn");
        var input9 =document.getElementById("jornadahn");
        var input10 =document.getElementById("nivelhn");
        var input11 =document.getElementById("ciclohn");
        var input12 =document.getElementById("zonahn");
      
        if (select.value === "14") {
          hiddenDiv.classList.remove("d-none");
      
          input1.required = true;
          input2.required = true;
          input3.required = true;
          input4.required = true;
          input5.required = true;
          input6.required = true;
          input7.required = true;
          input8.required = true;
          input9.required = true;
          input10.required = true;
          input11.required = true;
          input12.required = true;
      
        }
        else {
          hiddenDiv.classList.add("d-none");
      
          input1.required = false;
          input2.required = false;
          input3.required = false;
          input4.required = false;
          input5.required = false;
          input6.required = false;
          input7.required = false;
          input8.required = false;
          input9.required = false;
          input10.required = false;
          input11.required = false;
          input12.required = false;
        }
      }

      function actualizarCiudadesHN(paisSeleccionado) {
        var ciudadSelect = document.getElementById("munic");
        ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
      
        // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
        var xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function () {
          if (this.readyState == 4 && this.status == 200) {
            var opcionesCiudades = this.responseText;
            ciudadSelect.innerHTML = opcionesCiudades; // Actualizar la lista de ciudades
          }
        };
        xhr.open("POST", "/obtenerciudades", true);
        xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhr.send("pais=" + paisSeleccionado);
      }