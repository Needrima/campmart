let cartLinks = document.querySelectorAll(".cart-link")
// console.log(cartLink.getAttribute('value'))

cartLinks.forEach(link => {
    link.addEventListener('click', () => {
        let xhr = new XMLHttpRequest();

        xhr.open("POST", "/add-to-cart", true);

        console.log(link.getAttribute('value'))

        xhr.send(link.getAttribute('value')); 

        xhr.onreadystatechange = function() {
            if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
                alert(xhr.responseText);
            };
        };   
        
    })
})