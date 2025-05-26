// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
    'use strict'
  
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation');
    const hiddenInputs = document.querySelectorAll("form div[style*='display:none'] input");
  
    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
      form.addEventListener('submit', event => {
        requiredInputs.forEach(input => {
          input.removeAttribute("required");
          input.value = "";
        });
        if (!form.checkValidity()) {
          event.preventDefault()
          event.stopPropagation()
        }
  
        form.classList.add('was-validated')
      }, false)
    })
  })()
  
  
  
  
  function activaselect2(){
    var select = document.getElementById("TdC");
  var hiddenDiv = document.getElementById("hidden-question-div");
  var hiddenDiv2 = document.getElementById("hidden-question-div2");
  var hiddenDiv4 = document.getElementById("hidden-question-div5");
  var input = document.getElementById("UT");
  var input2 = document.getElementById("GA");
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
        hiddenDiv4.classList.add("d-none");
    input4.required = false;
    }
  
    }
  
  function activaselect3(){
  var hiddenDiv4 = document.getElementById("hidden-question-div5");
  var input2 = document.getElementById("GA");
  var input3 = document.getElementById("Turno");
    if(input2.value==="38"||input2.value==="39"||input2.value==="40"){ 
    input3.required = false;
        hiddenDiv4.classList.add("d-none");
      }else {
      input3.required = true;
      hiddenDiv4.classList.remove("d-none");
    }
  }
  
  function activaselect4(){
  var select = document.getElementById("FechaN");
  var hiddenDiv = document.getElementById("ocultar");
  var hiddenDiv2 = document.getElementById("menor");
  var hiddenDiv3 = document.getElementById("autorizacion");
  var hiddenDiv4 = document.getElementById("compartir_datos");
  var hiddenDiv5 = document.getElementById("banco");
  var hiddenDiv6 = document.getElementById("datos");
  var hiddenDiv7 = document.getElementById("usovoz");
  var birthDate = new Date(select.value);
  var diff_ms = new Date("2024-06-29") -birthDate ;
  var diff_ms2 = new Date("2024-07-06") -birthDate ;
  var age_dt = new Date(diff_ms); 
  var age_dt2 = new Date(diff_ms2); 
  var age= Math.abs(age_dt.getUTCFullYear() - 1970);
  var age2= Math.abs(age_dt2.getUTCFullYear() - 1970);
  var input = document.getElementById("docA");
  if(age2>=15 && age<22){    
    hiddenDiv.classList.remove("d-none");
        if(age>=18){
    hiddenDiv2.classList.remove("d-none");
    input.required = true;
    hiddenDiv3.classList.add("d-none");
    hiddenDiv4.classList.add("d-none");
    hiddenDiv5.classList.add("d-none");
    hiddenDiv6.classList.add("d-none");
    hiddenDiv7.classList.add("d-none");
  }
    else {
    hiddenDiv2.classList.add("d-none");
    hiddenDiv3.classList.remove("d-none");
    hiddenDiv4.classList.remove("d-none");
    hiddenDiv5.classList.remove("d-none");
    hiddenDiv6.classList.remove("d-none");
    hiddenDiv7.classList.remove("d-none");
    input.required = false;
  }
    } else {
    hiddenDiv.classList.add("d-none");
    }
    }
    
      function activaselect7(){
      var select = document.getElementById("gp");
      var hiddenDiv = document.getElementById("gps");
      if(select.value==="1"){
        hiddenDiv.classList.remove("d-none") ;
      }
      else {
         hiddenDiv.classList.add("d-none") ;
      }
      }
      
   
  
  
  function activaselect10(){
    var select = document.getElementById("dis");
  var hiddenDiv = document.getElementById("hidden-question-div10");
  var input4 = document.getElementById("dis2");
    if(select.value==="1"){    
    hiddenDiv.classList.remove("d-none");
    } else {
    hiddenDiv.classList.add("d-none");
      }
    }
  
  
    function validateFormat() {
      var formatInput = document.getElementById("DUIM").value;
         var formatRegex = /^\d{13}$/;
         var select = document.getElementById("nacionalidad");
        if(select.value==="1"||select.value===""){
              if (!formatRegex.test(formatInput)) {
            var alertContainer = document.getElementById("alertContainer");
            alertContainer.innerHTML = '<div class="alert alert-danger" role="alert">El formato del documento de identidad no es válido.</div>';
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
       function hideAlert2() {
      var alertContainer = document.getElementById("alertContainer2");
      alertContainer.style.display = "none";
    }
  
  function activaselect51(){
  var select = document.getElementById("nacionalidad");
  var labelOutput = document.getElementById("DUIMl");
  if(select.value==="1"){    
    labelOutput.innerHTML = "Escribe el número de tu Documento de Identidad"
    document.getElementById("DUIM").placeholder = "0000000000000";
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
  var input14= document.getElementById("direccion");
  var input19= document.getElementById("TM");
  var input20= document.getElementById("WS");
  var input22= document.getElementById("EstadoC");
  var input23= document.getElementById("TdC");
  var input28= document.getElementById("dis");
  var input30= document.getElementById("gp");
  var input33= document.getElementById("contacto");
  var input34= document.getElementById("contacto2");
  var input35= document.getElementById("relacion");
  var input37= document.getElementById("Transp");
  var input38= document.getElementById("gasto");
  var input39= document.getElementById("cantidad");
  var input40= document.getElementById("hijos");
  var input45= document.getElementById("docF");
  var input46= document.getElementById("docA");
  
  
  if(select.value==="1"){
    hiddenDiv.classList.remove("d-none") ;
    hiddenDiv2.classList.remove("d-none") ;
    hiddenDiv3.classList.add("d-none") 
   
    input9.required = true;
  input10.required = true;
  input11.required = true;
  input12.required = true;
  input13.required = true;
  input14.required = true;
  input19.required = true;
  input20.required = true;
  input22.required = true;
  input23.required = true;
  input28.required = true;
  input30.required = true;
  input33.required = true;
  input34.required = true;
  input35.required = true;
  input37.required = true;
  input38.required = true;
  input40.required = true;
  input45.required = true;
  input46.required = true;
  
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
     input14.required = false;
     input19.required = false;
     input20.required = false;
     input22.required = false;
     input23.required = false;
     input28.required = false;
     input30.required = false;
     input33.required = false;
     input34.required = false;
     input35.required = false;
     input37.required = false;
     input38.required = false;
     input39.required = false;
     input40.required = false;
     input45.required = false;
    input46.required = false;  
  }
  }
  
  function activa1(){
    var input = document.getElementById("nombreCompleto");
    var input2 = document.getElementById("apellido");
    var output = document.getElementById("nombrep");
    var output2 = document.getElementById("nombrep2");
      output.innerHTML = input.value +" "+input2.value;
      output2.innerHTML = input.value +" "+input2.value;
    
    }
  
    function actualizarCiudades(paisSeleccionado) {
      var ciudadSelect = document.getElementById("municipio");
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
  
   function actualizarCiudades2(paisSeleccionado) {
  var hiddenDiv = document.getElementById("Honduras");
  var hiddenDiv2 = document.getElementById("Guatemala");
  var select = document.getElementById("pais");
    if(select.value==="3"){
      hiddenDiv2.classList.add("d-none") ;
      hiddenDiv.classList.remove("d-none") ;
    }
    if(select.value==="4"){
      hiddenDiv2.classList.remove("d-none") ;
      hiddenDiv.classList.add("d-none") ;
    }
    var ciudadSelect = document.getElementById("departamento");
    var ciudadSelect2 = document.getElementById("pais2");
    ciudadSelect.innerHTML = "";
    ciudadSelect2.innerHTML = ""; // Limpiar la lista de ciudades
  
    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesCiudades = this.responseText;
            ciudadSelect.innerHTML = opcionesCiudades;
            ciudadSelect2.innerHTML = opcionesCiudades; // Actualizar la lista de ciudades
        }
    };
    xhr.open("POST", "/obtenerciudades2", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("pais=" + paisSeleccionado);
  
  }
  
  function actualizarCiudades3(paisSeleccionado) {
    var select = document.getElementById("pais");
    if(select.value==="4"||select.value==="3"){
    var ciudadSelect = document.getElementById("sede2");
    ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
  
    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesCiudades = this.responseText;
            ciudadSelect.innerHTML = opcionesCiudades; // Actualizar la lista de ciudades
        }
    };
    xhr.open("POST", "/obtenerciudades3", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("pais=" + paisSeleccionado);
  }
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
  
  function actualizarmun2(paisSeleccionado) {
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
    xhr.open("POST", "/obtenerciudades", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("pais=" + paisSeleccionado);
  }