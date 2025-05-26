
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
           else if (document.getElementById("alertContainer").style.display === "block") {
        event.preventDefault();
        alert("No se puede enviar el formulario mientras este incorrecto el DUI");
      }
        }, false)
      })
    })()
      
    function certifrenas() {
      var select = document.getElementById("FechaN");
      var select2 = document.getElementById("gi");
      var input = document.getElementById("fechaV");
      var input2= document.getElementById("docR");
    
      var birthDate = new Date(select.value);
      var diff_ms = Date.now() - birthDate;
      var age_dt = new Date(diff_ms);
      var age = Math.abs(age_dt.getUTCFullYear() - 1970);
    
      if(age >= 18) {    
        if(select2.value === "1") {
          input.required = true;
          input2.required = true;
          // código adicional aquí si select2.value es igual a "1"
        } else {
          input.required = false;
          input2.required = false;
          // código adicional aquí si select2.value no es igual a "1"
        }
      } else {
        input.required = false;
        input2.required = false;
      }
    }    
    
    function activaselect3(){
    var select = document.getElementById("TdC");
    var hiddenDiv = document.getElementById("hidden-question-div");
    var hiddenDiv2 = document.getElementById("hidden-question-div2");
    var hiddenDiv3 = document.getElementById("hidden-question-div3");
    var hiddenDiv4 = document.getElementById("hidden-question-div5");
    var input = document.getElementById("UT");
    var input2 = document.getElementById("GA");
    var input3 = document.getElementById("carrera");
    var input4 = document.getElementById("U");
    var input5 = document.getElementById("HS");
      if(select.value==="1"){    
          hiddenDiv.classList.add("d-none");
          input.required = false;
          hiddenDiv2.classList.remove("d-none");
          input2.required = true;
          input5.required = true;
          activaselect50();
          hiddenDiv4.classList.remove("d-none");
          
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
      input5.required = false;
      }
    
      }
    
    function activaselect50(){
    var hiddenDiv3 = document.getElementById("hidden-question-div3");
    var input2 = document.getElementById("GA");
    var input3 = document.getElementById("carrera");
    var input4 = document.getElementById("U");
      if(input2.value==="38"||input2.value==="39"||input2.value==="40"){ 
      hiddenDiv3.classList.remove("d-none");
      input3.required = true;
      input4.required = true;}
      else{
      hiddenDiv3.classList.add("d-none");
      input3.required = false;
      input4.required = false;
      }
      }
    
    function activaselect4(){
    var select = document.getElementById("FechaN");
    var hiddenDiv = document.getElementById("menor");
    var hiddenDiv2 = document.getElementById("mayor");
    var hiddenDiv3 = document.getElementById("divrenas");
    var input = document.getElementById("nombreRep");
    var input2 = document.getElementById("DUIP");
    var input3 = document.getElementById("parentesco");
    var input4 = document.getElementById("nombreAdolescente");
    var input5 = document.getElementById("edad");
    var input6 = document.getElementById("DUI");
    // var input7 = document.getElementById("fechaV");
    // var input8= document.getElementById("docR");
    
    var birthDate = new Date(select.value);
        var diff_ms = Date.now() -birthDate ;
        var age_dt = new Date(diff_ms); 
        var age= Math.abs(age_dt.getUTCFullYear() - 1970);
      if(age>=18){    
      hiddenDiv.classList.add("d-none");
      hiddenDiv2.classList.remove("d-none");
      hiddenDiv3.classList.remove("d-none");
      input.required=false;
      input2.required=false;
      input3.required=false;
      input4.required=false;
      input5.required=false;
      input6.required=true;
      // input7.required=true;
      // input8.required=true;  
      } else {
      hiddenDiv.classList.remove("d-none");
      hiddenDiv2.classList.add("d-none");
      hiddenDiv3.classList.remove("d-none");
      input.required=true;
      input2.required=true;
      input3.required=true;
      input4.required=true;
      input5.required=true;
      input6.required=false;
      // input7.required=false;
      // input8.required=false;
    }
      }
    
    function activainst(){
      var select =document.getElementById("acc2");
      var hiddenDiv = document.getElementById("divaccinst");
      var input = document.getElementById("accinst");
      if(select.value==="17"){
        hiddenDiv.classList.remove("d-none");
        input.required = true;
      }else{
        hiddenDiv.classList.add("d-none");  
        input.required = false;
      }
    }

    function activaselect30(){
    var select = document.getElementById("acc");
    var hiddenDiv = document.getElementById("acc2d");
    var input4 = document.getElementById("acc2");
    var hiddenDiv2 = document.getElementById("acc3d");
    var input5 = document.getElementById("acc3");
    var hiddenDiv3 = document.getElementById("puestod");
    var input6 = document.getElementById("puesto"); 
    
    if(select.value==="9"){    
      hiddenDiv.classList.remove("d-none");
      hiddenDiv2.classList.add("d-none");
      hiddenDiv3.classList.add("d-none"); 
      input4.required = true;
       input5.required = false;
       input6.required = false;
    
      }else if(select.value==="5"){    
      hiddenDiv3.classList.remove("d-none"); 
      hiddenDiv2.classList.remove("d-none");
      hiddenDiv.classList.add("d-none");
      input5.required = true;
     input6.required = true; 
      input4.required = false;
    } 
      else {
      hiddenDiv.classList.add("d-none");
      hiddenDiv2.classList.add("d-none");
      hiddenDiv3.classList.add("d-none");
      input4.required = false;
      input5.required = false;
      input6.required = false;
        }
      
      }
    
      
     
      
    function activaselect51(){
    var select = document.getElementById("gi");
    var labelOutput = document.getElementById("DUIL");
    var labelOutput2 = document.getElementById("Docl");
    var labelOutput3 = document.getElementById("DocAl");
    var hiddenDiv = document.getElementById("ocultar");
    var hiddenDiv2 = document.getElementById("ext2");
    var hiddenDiv3 = document.getElementById("ext");
    var hiddenDiv4 = document.getElementById("nacional");
    var hiddenDiv5 = document.getElementById("divrenas");
    var input= document.getElementById("pas");
    var input2= document.getElementById("docF");
    var input3= document.getElementById("docA");
    var input4= document.getElementById("paisn");
    // var input5= document.getElementById("docR");
    // var input6= document.getElementById("fechaV");
    hiddenDiv.classList.remove("d-none")
    if(select.value==="1"){    
      labelOutput.innerHTML ="Escribe el número de tu Documento Personal de Identidad (DPI)";
      labelOutput2.innerHTML ="Toma una fotografía a tu documento personal de identidad [DPI] (lado de enfrente) y súbelo en este espacio";
      labelOutput3.innerHTML ="Toma una fotografía a tu documento personal de identidad [DPI] (lado de atras) y súbelo en este espacio";
      document.getElementById("DUI").placeholder = "";
      document.getElementById("DUIP").placeholder = "";
      hiddenDiv2.classList.remove("d-none");
      hiddenDiv4.classList.remove("d-none");
      hiddenDiv3.classList.add("d-none");
      hiddenDiv5.classList.remove("d-none");
        input.required=false;
        input2.required=true;
        input3.required=true;
        input4.required=false;
        // input5.required=true;
        // input6.required=true;
        } else {
        
        labelOutput2.innerHTML ="Toma una fotografía a tu documento de identidad (lado de enfrente) y súbelo en este espacio";
        labelOutput3.innerHTML ="Toma una fotografía a tu documento de identidad (lado de atras) y súbelo en este espacio";
        document.getElementById("DUI").placeholder = "";
        document.getElementById("DUIP").placeholder = "";
        hiddenDiv3.classList.remove("d-none");
        hiddenDiv2.classList.add("d-none")
        hiddenDiv4.classList.add("d-none");
        hiddenDiv5.classList.add("d-none");
        input4.required=true;
        // input5.required=false;
        // input6.required=false;
      } 
      } 
    
    function activaselect52(){
    var select2 = document.getElementById("elec");
    var labelOutput = document.getElementById("DUIL");
    var hiddenDiv5 = document.getElementById("ext2");
    var hiddenDiv4 = document.getElementById("pasd");
    var hiddenDiv6 = document.getElementById("nacional");
    var input= document.getElementById("pas");
    var input2= document.getElementById("docF");
    var input3= document.getElementById("docA");
    hiddenDiv6.classList.remove("d-none");
    if(select2.value==="1"){
      labelOutput.innerHTML ="Escribe el número de tu Pasaporte";
        hiddenDiv4.classList.remove("d-none");
        hiddenDiv5.classList.add("d-none");
        input.required=true;
        input2.required=false;
        input3.required=false;
    } else if (select2.value==="2"){
      labelOutput.innerHTML ="Escribe el número de tu Documento de Identidad"
          hiddenDiv4.classList.add("d-none");
        hiddenDiv5.classList.remove("d-none");
        input.required=false;
        input2.required=true;
        input3.required=true;}
    else {
      hiddenDiv5.classList.add("d-none");
      hiddenDiv4.classList.add("d-none");
      hiddenDiv6.classList.add("d-none");
    }
    }
      
    function otro(){
      var select=document.getElementById('actividad');
      var select2=document.getElementById('subact');
      var hiddenDiv=document.getElementById('aD');
      if(select.value==="10"||select.value==="17"||select2.value==="17"||select2.value==="22"){
      hiddenDiv.classList.remove("d-none");
    }
    else {
      hiddenDiv.classList.add("d-none");
    }
    }
    
    function otro2(){
      var select=document.getElementById('OH');
      var hiddenDiv=document.getElementById('OHD');
    
      if(select.checked){
      hiddenDiv.classList.remove("d-none");
    }
    else {
      hiddenDiv.classList.add("d-none");
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

function actualizaractividades(areaSeleccionado) {
    var paisSeleccionado = document.getElementById("idp").value;
    var actividadSelect = document.getElementById("actividad");
    actividadSelect.innerHTML = ""; // Limpiar la lista de actividades

    // Enviar petición AJAX para obtener las actividades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesactividad = this.responseText;
            actividadSelect.innerHTML = opcionesactividad; // Actualizar la lista de actividades
        }
    };
    xhr.open("POST", "/obteneractividad", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("area=" + areaSeleccionado +"&pais="+paisSeleccionado);

    var depSelect = document.getElementById("departamento2");
    depSelect.innerHTML = ""; // Limpiar la lista de actividades

    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesdep2 = this.responseText;
            depSelect.innerHTML = opcionesdep2; // Actualizar la lista de actividades
        }
    };
    xhr.open("POST", "/obtenerdepartamento2", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("area=" + areaSeleccionado +"&pais="+paisSeleccionado);
}

function actualizarsubactividades(actividadSeleccionado) {
      var select2=document.getElementById('actividad');
      var pais=document.getElementById('idp').value;
    var hiddenDiv=document.getElementById('subactD');
    if(select2.value==="1"||select2.value==="4"){
      hiddenDiv.classList.remove("d-none");
    }
    else {
      hiddenDiv.classList.add("d-none"); 
    }
  var subactividadSelect = document.getElementById("subact");
    subactividadSelect.innerHTML = ""; // Limpiar la lista de actividades

    // Enviar petición AJAX para obtener las actividades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionessubactividad = this.responseText;
            subactividadSelect.innerHTML = opcionessubactividad; // Actualizar la lista de actividades
        }
    };
    xhr.open("POST", "/obtenersubactividad", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("actividad=" + actividadSeleccionado  +"&pais="+pais);
  }

  
  function actualizarsedes2(actividadSeleccionado) {
  var paisSeleccionado = document.getElementById("pais2").value;
  var subactividadSelect = document.getElementById("sede2");
  subactividadSelect.innerHTML = ""; // Limpiar la lista de actividades
  // Enviar petición AJAX para obtener las actividades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var opcionessubactividad = this.responseText;
          subactividadSelect.innerHTML = opcionessubactividad; // Actualizar la lista de actividades
      }
  };
  xhr.open("POST", "/obtenersede2", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("actividad=" + actividadSeleccionado +"&area="+paisSeleccionado);
}

