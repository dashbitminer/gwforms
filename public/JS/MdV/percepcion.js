// Ejemplo de JavaScript inicial para deshabilitar el envío de formularios si hay campos no válidos
(function () {
    'use strict'
  
    // Obtener todos los formularios a los que queremos aplicar estilos de validación de Bootstrap personalizados
    var forms = document.querySelectorAll('.needs-validation')
  
    // Bucle sobre ellos y evitar el envío
    Array.prototype.slice.call(forms)
      .forEach(function (form) {
        form.addEventListener('submit', function (event) {
          if (!form.checkValidity()) {
            event.preventDefault()
            event.stopPropagation()
          }
  
          form.classList.add('was-validated')
        }, false)
      })
  })()


function participarF() {
    var select = document.getElementById("participar");
    var hiddenDiv = document.getElementById("ocultar");
    var hiddenDiv2 = document.getElementById("botondiv");
    hiddenDiv2.classList.remove("d-none");
var input3=document.getElementById("edad");
var input4=document.getElementById("Sexo");
var input5=document.getElementById("p1");
var input6=document.getElementById("p2");
var input8=document.getElementById("p3");
var input9=document.getElementById("p31");
var input10=document.getElementById("p4");
var input11=document.getElementById("p41");
var input12=document.getElementById("p5");
var input13=document.getElementById("p51");
var input14=document.getElementById("p6");
var input15=document.getElementById("p61");
var input16=document.getElementById("p7");
var input17=document.getElementById("p8");
var input18=document.getElementById("p9");
var input19=document.getElementById("p91");
var input20=document.getElementById("p10");
var input21=document.getElementById("p101");
var input22=document.getElementById("p11");
var input23=document.getElementById("p12");

  if(select.value==="1"){
    hiddenDiv.classList.remove("d-none");
    input3.required = true
input4.required = true
input5.required = true
input6.required = true
input8.required = true
input9.required = true
input10.required = true
input11.required = true
input12.required = true
input13.required = true
input14.required = true
input15.required = true
input16.required = true
input17.required = true
input18.required = true
input19.required = true
input20.required = true
input21.required = true
input22.required = true
input23.required = true
  }
  else {
    hiddenDiv.classList.add("d-none");
    input3.required = false
input4.required = false
input5.required = false
input6.required = false
input8.required = false
input9.required = false
input10.required = false
input11.required = false
input12.required = false
input13.required = false
input14.required = false
input15.required = false
input16.required = false
input17.required = false
input18.required = false
input19.required = false
input20.required = false
input21.required = false
input22.required = false
input23.required = false
  }
  }

  function p2F() {
    var hiddenDiv = document.getElementById("p2div");
    var select = document.getElementById("p2");
    var input7=document.getElementById("p21");
    if(select.value==="1"){
    input7.required = true
    hiddenDiv.classList.remove("d-none");
    }
    else{
      input7.required = false
      hiddenDiv.classList.add("d-none");
    }
  }