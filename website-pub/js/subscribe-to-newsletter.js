let subscriptionBtn = document.querySelector('#subscription-button')

subscriptionBtn.addEventListener('click', () => {
    let email = document.querySelector('#newsletter-form');
    console.log(email.value)

    let xhr = new XMLHttpRequest();    

    xhr.open("POST", "/subscribe-to-newsletter", true);    

    xhr.send(email.value); 

    xhr.onreadystatechange = function() {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            alert(xhr.responseText);
        };
    }; 

    email.value = '';
})
