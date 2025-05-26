function buscarm(valor) {
    var ciudadSelect = document.getElementById("municipioResidencia");
    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesmunsede= this.responseText;
            ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
            ciudadSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
        }
    };
    xhr.open("POST", "/obtenermunicipiosql", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("departamento=" + valor);
  }
  
  function guardar(event){
   var form = document.getElementById("form");
  var invalidElements = form.querySelectorAll('.is-invalid');
   event.preventDefault();
   if (!form.checkValidity() || invalidElements.length > 0) {
    event.preventDefault();
    event.stopPropagation();
    Swal.fire({
      icon: 'error',
      title: 'Oops...',
      text: 'Por favor, completa el formulario correctamente antes de enviar.',
    });
    return;
  }
    
    var formData = new FormData(form);
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
        var respuesta = this.responseText;
       if(respuesta == "1"){
    Swal.fire({
      title: 'Participante registrado correctamente',
      icon: 'success'
    }) .then((result) => {
            // Redirige a la página de preinscripción después de que el usuario cierre el SweetAlert
            window.location.href = "/juventud/cp/preinscripcion";
          });
        } else {
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Error al registrar el participante',
          }).then((result) => {
          });
        }
      }
    };
    xhr.open("POST", "/juventud/cp/preinscripcion/post", true);
    xhr.send(formData);
  }
  
  document.getElementById("form").addEventListener('submit', guardar);
  
  var formElements = form.querySelectorAll('input, textarea, select');
  for (var i = 0; i < formElements.length; i++) {
    if (formElements[i].name !== 'fechaNacimiento') {
      formElements[i].addEventListener('change', function() {
        if (this.validity.valid) {
          this.classList.remove('is-invalid');
          this.classList.add('is-valid');
        } else {
          this.classList.remove('is-valid');
          this.classList.add('is-invalid');
        }
      });
    }
  }
  
  function comprobarfecha() {
    var fechaNacimiento = document.getElementById("fechaNacimiento");
    var hidden = document.getElementById("alertaFecha");
    var hoy = new Date();
    var cumpleanos = new Date(fechaNacimiento.value);
    var edad = hoy.getFullYear() - cumpleanos.getFullYear();
    var m = hoy.getMonth() - cumpleanos.getMonth();
    if (m < 0 || (m === 0 && hoy.getDate() < cumpleanos.getDate())) {
      edad--;
    }
    if(edad < 18 && edad >= 15){
      fechaNacimiento.classList.remove('is-invalid');
    fechaNacimiento.classList.add('is-valid');
    hidden.classList.remove('d-none');
      document.getElementById("menor").classList.remove('d-none');
      document.getElementById("nombreRep").required = true;
      document.getElementById("telefonoRep").required = true;
    }else if(edad < 15 || edad > 29){    
  Swal.fire({
    text: 'Agradecemos tu interés en participar en el programa Creando Profesionales. Lastimosamente el perfil que estamos buscando para esta actividad es de jóvenes de 15 a 29 años. Sin embargo, te invitamos a estar pendiente de nuestras redes sociales en Instagram: Glasswing El Salvador (glasswingsv) y Parque Cuscatlán (miparquecuscatlán) y en Facebook: Glasswing El Salvador para que conozcas otras actividades en las que nos encantaría contar con tu participación.',
  icon: 'info',
  customClass: {
    content: 'text-justify'
  }
  }) .then((result) => {
    fechaNacimiento.classList.remove('is-valid');
    fechaNacimiento.classList.add('is-invalid');
    hidden.classList.add('d-none');
      });
      
  }else{     
    fechaNacimiento.classList.remove('is-invalid');
    fechaNacimiento.classList.add('is-valid');
    hidden.classList.remove('d-none');
    document.getElementById("menor").classList.add('d-none');
    document.getElementById("nombreRep").required = false;
    document.getElementById("telefonoRep").required = false;
  }}
  
  window.onload = function() {
    var isMobile = /iPhone|iPad|iPod|Android/i.test(navigator.userAgent);
    var fechaNacimiento = document.getElementById('fechaNacimiento');
  
    if (isMobile) {
      fechaNacimiento.onchange = comprobarfecha;
    }
  };
  
  function mostrar(valor){
    var hidden = document.getElementById("todo");
    if(valor == "1"){
      hidden.classList.remove('d-none');
    }else{
      hidden.classList.add('d-none');
    }
  }