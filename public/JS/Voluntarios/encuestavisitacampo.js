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

  function mostraobservacion(selectId, divId){
    var select = document.getElementById(selectId);
    var hiddenDiv = document.getElementById(divId);
      if(select.value===""){
          hiddenDiv.classList.add("d-none"); 
      }
      else{
          hiddenDiv.classList.remove("d-none");
      }
  }

  function actualizaractividades(areaSeleccionado) {
   
    var actividadSelect = document.getElementById("p86");
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
    xhr.send("area=" + areaSeleccionado +"&pais=-1");
    }

    function actualizarmun() {
      var programaselect = document.getElementById("p158").value;
      var paisSeleccionado = document.getElementById("p188").value;
      var ciudadSelect = document.getElementById("p84");
      $('#p84').selectize()[0].selectize.destroy(); 
      ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
      // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
      var xhr = new XMLHttpRequest();
      xhr.onreadystatechange = function() {
          if (this.readyState == 4 && this.status == 200) {
              var opcionesmunsede= this.responseText;
              ciudadSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
              $('#p84').selectize({
                sortField: 'text'
            });
            }
      };
      xhr.open("POST", "/obtenersedes", true);
      xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
      xhr.send("departamento=" + paisSeleccionado+"&programa="+programaselect);
      ciudadSelect.required = true;
    }
    // 

    function actualizardepto(paisSeleccionado) {
      var ciudadSelect = document.getElementById("p188");
      $('#p188').selectize()[0].selectize.destroy(); 
      ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
      // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
      var xhr = new XMLHttpRequest();
      xhr.onreadystatechange = function() {
          if (this.readyState == 4 && this.status == 200) {
              var opcionesmunsede= this.responseText;
              ciudadSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
              $('#p188').selectize({
                sortField: 'text'
            });
            }
      };
      xhr.open("POST", "/obtenerdepartamentog", true);
      xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
      xhr.send("pais=" + paisSeleccionado);
      ciudadSelect.required = true;
    }

    function ocultardiv(){
      
        var input= document.getElementById("actividad").value;
        var div1= document.getElementById("presencial1");
        var input1=document.getElementById("p116")
        var input2=document.getElementById("p117")
        var input3=document.getElementById("p118")

        var div2=document.getElementById("virtual");
        var input4=document.getElementById("p119")
        var input5=document.getElementById("p112")
        var input6=document.getElementById("p107")
        var input7=document.getElementById("p121")
        var input8=document.getElementById("p122")
        var input9=document.getElementById("p123")

        var div3=document.getElementById("presencial2");
        var input10=document.getElementById("p124")
        var input11=document.getElementById("p127")

        if (input==="1"){
        div1.classList.remove("d-none");
        input1.required = true;
        input2.required = true;
        input3.required = true;

        div2.classList.add("d-none")
        input4.required = false;
        input5.required = false;
        input6.required = false;
        input7.required = false;
        input8.required = false;
        input9.required = false;
        
        div3.classList.remove("d-none");
        input10.required = true;
        input11.required = true;

    }else{
      div1.classList.add("d-none");
      input1.required = false;
      input2.required = false;
      input3.required = false;

      div2.classList.remove("d-none")
      input4.required = true;
      input5.required = true;
      input6.required = true;
      input7.required = true;
      input8.required = true;
      input9.required = true;

      div3.classList.add("d-none");
      input10.required = false;
      input11.required = false;

    }

    }