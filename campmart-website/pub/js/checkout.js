function checkCheckout() {
    let buyerName = document.querySelector('#buyerName');
    let buyerEmail = document.querySelector('#buyerEmail');
    let buyerPhone = document.querySelector('#buyerPhone');
    let comment = document.querySelector('#comment');

    let inputValues = [buyerName, buyerEmail, buyerPhone];
    inputValues.forEach(val => {
        if (val.value.trim().length < 3) {
            alert('Please fill in the required fields with atleast three characters*');
            return false;
        }

        return true;
    })
}