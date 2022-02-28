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

//shop.html only
let searchInput = document.querySelector('#searchInput');

function checkSearch() {
    if (searchInput.value === '') {
        alert('You did not enter a value to search');
        return false;
    }

    return true;
}
//shop.html only end

//contact.html only

function checkContactMsg() {
    let senderName = document.querySelector('#senderName');
    let senderEmail = document.querySelector('#senderEmail');
    let msg = document.querySelector('#msg');
    let msgSubject = document.querySelector('#msgSubject');

    let inputValues = [senderName, senderEmail, msg, msgSubject];
    inputValues.forEach(val => {
        if (val.value === '') {
            alert('Please fill in the required fields marked *');
            return false;
        }

        return true;
    })
}
//contact.html only end

// single-product.html only

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

//single-product.html only


// cart.html only

// qty, prices and subtotals have the same length
//which is number of items in the cart
let qtyInput = document.querySelectorAll('.qty');

function initValues() {
    let prices = document.querySelectorAll('.price'); // item price
    let subtotals = document.querySelectorAll('.s-total'); // item subtotal  = price * qty
    
    let cartTotal = document.querySelector('#cartTotal') // cart total
    let shippingFee = document.querySelector('#shipping-fee').textContent; // shipping fee   
    let orderTotal = document.querySelector('#orderTotal'); // order total = cart total + shipping fee
    
    let totalValue = 0; //initialized to hold cart total value

    subtotals.forEach((subtotal, index) => {
        //takout 'NGN' in front of price
        // get quantity
        // multiply price by qty and assign to a variable itemTotalValue
        let price = parseInt(prices[index].textContent.slice(3));
        let qty = qtyInput[index].value;                
        let itemTotalValue = price * qty;

        // set item subtotal to price * qty
        subtotal.textContent = 'NGN'+itemTotalValue.toString(); 

        // sum up item subtotal and assign to variable totalValue
        totalValue += itemTotalValue;
    })

    // set cart total value to "totalValue"
    cartTotal.textContent = 'NGN'+totalValue.toString();

    // remove first three value from shipping fee
    // shipping value will be a string with value 'e' if value is free
    // or shipping value if not free
    let shippingFeeValue = shippingFee.slice(3)

    let val = 0;// val to hold delivery fee

    // if delivery is not free
    //convert to value to number and set to variable val
    if (!isNaN(shippingFeeValue)) {
        val = parseInt(shippingFeeValue)
    }

    // set order value to totalValue + shipping fee
    let orderValue = totalValue + val;
    orderTotal.textContent = 'NGN'+orderValue.toString()
}
initValues();

// reset quantity to 1 if it is less than one and
// call initValues() whenever quantity is changed 
qtyInput.forEach(qty => {
    qty.addEventListener('input', () => {
        // alert('Hellow world')
        if (qty.value < 1) {
            qty.value = 1;
        }

        initValues();
    })
})

//cart.html only end
