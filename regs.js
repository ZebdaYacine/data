function padStart(value, length, padChar) {
    value = String(value); // Ensure it's a string
    while (value.length < length) {
        value = padChar + value;
    }
    return value;
}
var date=null
// Modify the original script
if (date_prestation) {
    date_prestation = date_prestation.replace(/[\\/]/g, '-');
    if (date_prestation.match(/^\d{4}-\d{1,2}-\d{1,2}$/)) {
        var parts = date_prestation.split("-");
        date = padStart(parts[2], 2, '0') + "-" + padStart(parts[1], 2, '0') + "-" + parts[0];
    } else if (date_prestation.match(/^\d{1,2}-\d{1,2}-\d{4}$/)) {
        var parts = date_prestation.split("-");
        date = padStart(parts[0], 2, '0') + "-" + padStart(parts[1], 2, '0') + "-" + parts[2];
    } else if (date_prestation.match(/^\d{4}\/\d{1,2}\/\d{1,2}$/)) {
        var parts = date_prestation.split("/");
        date = padStart(parts[2], 2, '0') + "-" + padStart(parts[1], 2, '0') + "-" + parts[0];
    } else if (date_prestation.match(/^\d{1,2}\/\d{1,2}\/\d{4}$/)) {
        var parts = date_prestation.split("/");
        date = padStart(parts[0], 2, '0') + "-" + padStart(parts[1], 2, '0') + "-" + parts[2];
    } else if (date_prestation.match(/^\d{4}\\\d{1,2}\\\d{1,2}$/)) {
        var parts = date_prestation.split("\\");
        date = padStart(parts[2], 2, '0') + "-" + padStart(parts[1], 2, '0') + "-" + parts[0];
    } else if (date_prestation.match(/^\d{1,2}\\\d{1,2}\\\d{4}$/)) {
        var parts = date_prestation.split("\\");
        date_achat = padStart(parts[0], 2, '0') + "-" + padStart(parts[1], 2, '0') + "-" + parts[2];
    } else {
        date = null;
    }
}

date;