
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