// all pages script
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

let newsletterForm = document.querySelector('#newsletter-form');

function newsletterSub() {
    if (newsletterForm.value === '') {
        alert('You did not input your E-mail!');
        return false;
    }

    return true;
}
// all pages script end


