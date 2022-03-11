let subscriptionBtn = document.querySelector('#subscription-button')

subscriptionBtn.addEventListener('click', () => {
    let email = document.querySelector('#newsletter-form').value;
    console.log(email)

    let xhr = new XMLHttpRequest();    

    xhr.open("POST", "/subscribe-to-newsletter", true);    

    xhr.send(email); 

    xhr.onreadystatechange = function() {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            alert(xhr.responseText);
        };
    }; 
})
