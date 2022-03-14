let cartBtn = document.querySelector("#cart-btn");

cartBtn.addEventListener('click', () => {
    let id = cartBtn.getAttribute("value");
    console.log(id)
    
    let qty = document.querySelector("#product-qty").value;
    console.log(qty)
    
    let type = document.querySelector("#types").value;
    console.log(type)

    let xhr = new XMLHttpRequest();

    xhr.open("POST", "/single-to-cart", true);

    console.log(id);

    let data = id+" "+qty+" "+type; // Ex. "675030nvjdkshg84ndj 3 small"

    xhr.send(data); 

    xhr.onreadystatechange = function() {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            alert(xhr.responseText);
        };
    };   
    
})