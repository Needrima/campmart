let searchInput = document.querySelector('#searchInput');

function checkSearch() {
    if (searchInput.value === '') {
        alert('You did not enter a value to search');
        return false;
    }

    return true;
}