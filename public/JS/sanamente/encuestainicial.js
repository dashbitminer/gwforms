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
      }, false)
    })
  })()

  var temperatureElement = document.querySelector('.temperature');

  function updateTemperature(temp) {
    temperatureElement.setAttribute('data-temp', temp);
    var maxHeight = 150; // Altura máxima del termómetro en píxeles
    var height = (temp / 10) * maxHeight; // Ajustar la altura según el valor de temperatura
    temperatureElement.style.height = height + "px";
  }

  function mostrar(){
    var select =document.getElementById('puesto')
    var hiddenDiv1 =document.getElementById('a')
    var hiddenDiv2 =document.getElementById('b')
    var hiddenDiv3 =document.getElementById('c')
    var hiddenDiv4 = document.getElementById('otro1')
    var hiddenDiv5 = document.getElementById('bEst')

    var input1 = document.getElementById('docente')
    var input2 = document.getElementById('ntrabaja')
    var input3 = document.getElementById('catedra')
    var input4 = document.getElementById('otrorol')
    var input5 = document.getElementById('ntrabajaEst')

    if (select.value==="3") {
        hiddenDiv1.classList.remove("d-none");
        hiddenDiv2.classList.remove("d-none");
        hiddenDiv3.classList.remove("d-none");
        hiddenDiv4.classList.add("d-none");
        hiddenDiv5.classList.add("d-none")

        input1.required = true;
        input2.required = true;
        input3.required = true;
        input4.required = false;
        input4.value = "";
        input5.required = false;
        input5.value = "";

    } else if (select.value==="6") {
      hiddenDiv1.classList.add("d-none");
      hiddenDiv2.classList.add("d-none");
      hiddenDiv3.classList.add("d-none");
      hiddenDiv4.classList.remove("d-none");
      hiddenDiv5.classList.add("d-none");

      input1.required = false;
      input2.required = false;
      input3.required = false;
      input1.value = "";
      input2.value = "";
      input3.value = "";
      input4.required = true;
      input5.required = false;
      input5.value = "";

    } else if (select.value==="5") {
      hiddenDiv1.classList.add("d-none");
      hiddenDiv2.classList.add("d-none");
      hiddenDiv3.classList.add("d-none");
      hiddenDiv4.classList.add("d-none");
      hiddenDiv5.classList.remove("d-none")

      input1.required = false;
      input2.required = false;
      input3.required = false;
      input1.value = "";
      input2.value = "";
      input3.value = "";
      input4.required = false;
      input4.value = "";
      input5.required = true;

    } else {
        hiddenDiv1.classList.add("d-none");
        hiddenDiv2.classList.add("d-none");
        hiddenDiv3.classList.add("d-none");
        hiddenDiv4.classList.add("d-none");
        hiddenDiv5.classList.add("d-none");

        input1.required = false;
        input2.required = false;
        input3.required = false;
        input4.required = false;
        input5.required = false;
        input1.value = "";
        input2.value = "";
        input3.value = "";
        input4.value = "";
        input5.value = "";
    }

  }

  function mostrarEF(){
    var select =document.getElementById('puesto')
    var hiddenDiv4 = document.getElementById('otro1')

    var input4 = document.getElementById('otrorol')

    if (select.value==="6") {
      hiddenDiv4.classList.remove("d-none");

      input4.required = true;

    } else {
        hiddenDiv4.classList.add("d-none");

        input4.value = "";
    }

  }

  function mostrar2(valor){
    var hiddenDiv1 =document.getElementById('ayuda')
    var checkbox = document.getElementsByName('ayuda')
    if(valor==="1"){ 
        hiddenDiv1.classList.add("d-none");
        checkbox.required = false;
    }else {
        hiddenDiv1.classList.remove("d-none");
        checkbox.required = false;
    }
  }

  function mostrar3(){
    var select =document.getElementById('puesto')
    var hiddenDiv1 =document.getElementById('question5')
    var input1 = document.getElementsByName('situacion31')
    var input2 = document.getElementsByName('situacion32')
    var input3 = document.getElementsByName('situacion33')
    var input4 = document.getElementsByName('situacion34')
    var input5 = document.getElementsByName('situacion35')
    var input6 = document.getElementsByName('situacion36')
    var input7 = document.getElementsByName('situacion37')
    if(select.value==="3"||select.value==="5"){ 
        hiddenDiv1.classList.remove("d-none");
        input1.required = true;
        input2.required = true;
        input3.required = true;
        input4.required = true;
        input5.required = true;
        input6.required = true;
        input7.required = true;
    }else {
        hiddenDiv1.classList.add("d-none");
        input1.required = false;
        input2.required = false;
        input3.required = false;
        input4.required = false;
        input5.required = false;
        input6.required = false;
        input7.required = false;
    }
  }

  function mostrar4(){
    var select =document.getElementById('puesto')
    var hiddenDiv1 =document.getElementById('question6')
    var input1 = document.getElementsByName('situacion31')
    var input2 = document.getElementsByName('situacion32')
    var input3 = document.getElementsByName('situacion33')
    var input4 = document.getElementsByName('situacion34')
    var input5 = document.getElementsByName('situacion35')
    var input6 = document.getElementsByName('situacion36')
    var input7 = document.getElementsByName('situacion37')
    if(select.value==="3"||select.value==="5"){ 
        hiddenDiv1.classList.remove("d-none");
        input1.required = true;
        input2.required = true;
        input3.required = true;
        input4.required = true;
        input5.required = true;
        input6.required = true;
        input7.required = true;
    }else {
        hiddenDiv1.classList.add("d-none");
        input1.required = false;
        input2.required = false;
        input3.required = false;
        input4.required = false;
        input5.required = false;
        input6.required = false;
        input7.required = false;
    }
  }

  document.addEventListener("DOMContentLoaded", function() {
    const images = document.querySelectorAll(".img-max");
  
    images.forEach(function(image) {
      image.addEventListener("click", function() {
        const radio = image.closest(".form-check").querySelector(".form-check-input");
        radio.checked = true;
      });
    });
  });

  document.getElementById('formualrioRegistro').addEventListener('submit', function(e) {
    var ayudaDiv = document.getElementById('ayuda');
    if (!ayudaDiv.classList.contains('d-none')) {
      var checkboxes = ayudaDiv.querySelectorAll('input[type="checkbox"]');
      var checked = Array.from(checkboxes).some(checkbox => checkbox.checked);
      if (!checked) {
        e.preventDefault();
        alert('Por favor, selecciona al menos una opción de ayuda.');
      }
    }
  });

  function acepta(){
    var hiddenDiv1 = document.getElementById('acepta')
    var input1 = document.getElementById('m1')
    if(input1.checked){    
      hiddenDiv1.classList.remove("d-none");
      } else {
      hiddenDiv1.classList.add("d-none");
      }
  }

  function autocuidado(valor){
    var hiddenDiv1 =document.getElementById('fomentaautoc1')
    var checkbox = document.getElementsByName('fomentaautoc')
    if(valor==="1"||valor==="2"){ 
        hiddenDiv1.classList.remove("d-none");
        checkbox.required = false;
    }else {
        hiddenDiv1.classList.add("d-none");
        checkbox.required = false;
    }
  }

  function consciente(valor){
    var hiddenDiv1 =document.getElementById('conocimientoT1')
    var input = document.getElementsByName('conocimientoT')
    if(valor==="1"){ 
        hiddenDiv1.classList.remove("d-none");
        input.required = true;
        input.value = "";
    }else {
        hiddenDiv1.classList.add("d-none");
        input.required = false;
    }
  }

  function referido(checkbox) {
    var hiddenDiv1 = document.getElementById('referido1');
    var input = document.getElementById('nreferido');
    
    // Comprobar si el checkbox con valor 6 está seleccionado
    var isChecked = document.querySelector('input[name="practicasAt"][value="6"]').checked;
    
    if (isChecked) {
      hiddenDiv1.classList.remove("d-none");
      input.required = true;
    } else {
      hiddenDiv1.classList.add("d-none");
      input.required = false;
      input.value = ""; // Clear input value when hiding
    }
  }

  function climaL(valor){
    var hiddenDiv1 =document.getElementById('climaT1')
    var input = document.getElementsByName('climaT')
    if(valor==="2"){ 
        hiddenDiv1.classList.remove("d-none");
        input.required = true;
        input.value = "";
    }else {
        hiddenDiv1.classList.add("d-none");
        input.required = false;
    }
  }