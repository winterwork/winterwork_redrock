let reg = document.getElementById('reg');
let login = document.getElementById('login');
let btn1 = document.getElementById('c1');
let btn2 = document.getElementById('c2');
const line1 = document.getElementById('line1');
const line2 = document.getElementById('line2');
const c1 = document.getElementById('c1');
const c2 = document.getElementById('c2');

c1.addEventListener('click', () => {
    reg.style.display = 'none';
    login.style.display = 'block';
    line1.style.backgroundColor = 'black';
    line2.style.backgroundColor = 'rgb(212, 212, 212)';
    c1.style.color = 'black';
    c2.style.color = 'rgb(212, 212, 212)';
})

c2.addEventListener('click', () => {
    reg.style.display = 'block';
    login.style.display = 'none';
    line2.style.backgroundColor = 'black';
    line1.style.backgroundColor = 'rgb(212, 212, 212)';
    c2.style.color = 'black';
    c1.style.color = 'rgb(212, 212, 212)';
})

