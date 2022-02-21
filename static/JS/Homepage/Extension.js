const extbt = document.getElementById('userHome')
const ext = document.getElementById('ext')
// console.log(extbt);
// console.log(ext);
ext.style.display = 'none';
let flag = true;
extbt.addEventListener('click', () => {
    if (flag) {
        ext.style.display = 'block';
        flag = false;
    }
    else {
        ext.style.display = 'none';
        flag = true;
    }
})