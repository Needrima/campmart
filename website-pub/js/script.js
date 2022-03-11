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

// all pages script end




