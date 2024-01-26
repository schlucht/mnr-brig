const URL = window.location.href

function priceValidate(price) {
    if (!price || !parseFloat(price)) return 0.0    
    return price
}

function stringIsValidate(txt) {
    return txt ? true : false
}



