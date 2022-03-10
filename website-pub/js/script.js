// all pages script
// navbar responsiveness
const bar = document.querySelector('#bar');
const nav = document.querySelector('#navbar');
const cl = document.querySelector('#close');

if (bar) {
    bar.addEventListener('click', () => {
        nav.classList.add('active');
    })
} 

if (cl) {
    cl.addEventListener('click', () => {
        nav.classList.remove('active');
    })
}

// subscription trial
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

// all pages script end




