let MainImg = document.getElementById('MainImg');
let smallimgs = document.getElementsByClassName('small-img');

for (let i = 0; i < smallimgs.length; i++) {
    smallimgs[i].addEventListener('click', () => {
        MainImg.src = smallimgs[i].src;
    })
}

let productQty = document.querySelector('#product-qty');

productQty.addEventListener('input', () => {
    if (productQty.value < 1) {
        productQty.value = 1;
    }
})

function addItemToCart() {
    let productTypes = document.querySelector('#types')
    if (productTypes.selectedIndex === 0) {
        alert('Please select a '+productTypes.options[0].value+' for your product')
        return false
    }

    return true
}