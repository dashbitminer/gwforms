/* // Example starter JavaScript for disabling form submissions if there are invalid fields
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
  })() */

function activaselect1(){
    var select = document.getElementById("p45");
    var hiddenDiv = document.getElementById("otrol");
    var input = document.getElementById("noasiste")

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

/* function activaselect2() {
  var hiddenDiv = document.getElementById('otrol2');
  var input = document.getElementById('otro2');

  // Comprobar si la opción con valor 42 está seleccionada en el select múltiple
  var isSelected = document.querySelector('select[name="p3"] option[value="18"]:checked');

  if (isSelected) {
    hiddenDiv.classList.remove("d-none");
    input.required = true;
  } else {
    hiddenDiv.classList.add("d-none");
    input.required = false;
    input.value = ""; // Limpiar el valor cuando se oculta el div
  }
} */

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

function activaselect3() {
  var select = document.getElementById("p51");
    var hiddenDiv = document.getElementById("otrol3");
    var input = document.getElementById("comentarios")

    if(select.value==="65"){
      input.required=true;
      hiddenDiv.classList.remove("d-none");
    }
    else{
      input.required=false;
      input.value="";
      hiddenDiv.classList.add("d-none");
    }
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
     title: 'Respuesta registrada correctamente',
     icon: 'success'
   }) .then((result) => {
           // Redirige a la página de preinscripción después de que el usuario cierre el SweetAlert
           window.location.href = "/voluntariado/encuesta/respuesta";
         });
       } else {
         Swal.fire({
           icon: 'error',
           title: 'Oops...',
           text: 'Error al registrar la respuesta',
         }).then((result) => {
         });
       }
     }
   };
   xhr.open("POST", "/voluntariado/encuesta/evento/post", true);
   xhr.send(formData);
 }
 
 document.getElementById("form").addEventListener('submit', guardar);
