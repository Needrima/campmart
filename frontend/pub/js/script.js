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

function newsletterSub(evt) {
    if (newsletterForm.value === '') {
        alert('You did not input your E-mail!');
        return false;
    }

    return true;
}
// cart.html 
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

//cart,html only end
window.onload = function() {
    $('.tc').click(function() {
        $('.tandc').css('display', 'block');

        $('.tandc').dialog({
            draggable: false,
            modal: true,
            resizable:true,
            height: 400,
            width: 300,

            buttons: [
                {
                    text: "I Agree",
                    icon: 'ui-icon-check',
                    click: function() {
                        $(this).dialog('close')
                    },
                },
            ]
        });
    })

    $('.pp').click(function() {
        $('.privacy').css('display', 'block');

        $('.privacy').dialog({
            draggable: false,
            modal: true,
            resizable:true,
            height: 300,
            width: 500,

            buttons: [
                {
                    text: "Ok",
                    icon: 'ui-icon-check',
                    click: function() {
                        $(this).dialog('close')
                    },
                },
            ]
        });
    })

    $('.di').click(function() {
        $('.delInfo').css('display', 'block');

        $('.delInfo').dialog({
            draggable: false,
            modal: true,
            resizable:true,
            height: 300,
            width: 300,

            buttons: [
                {
                    text: "Ok",
                    icon: 'ui-icon-check',
                    click: function() {
                        $(this).dialog('close')
                    },
                },
            ]
        });
    })
}