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