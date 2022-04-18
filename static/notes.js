const note = document.getElementById("text");
const shortcutKeys = {
    "s": function() {
        console.log("saved");
    },
    "b": function() {
        console.log(document.getSelection().getRangeAt(0));
    },
    "i": function() {

    },
    "u": function() {

    },
    "r": function() {
        var results = note.innerHTML.matchAll(/(\*{1,3})([^*]+)\1/g)
        var offset = 0;
        for (var result of results) {
            note.innerHTML = replaceRange(
                note.innerHTML,
                result.index + offset,
                result.index + result[0].length + offset,
                genReplaceString(result[1], result[2])
            );
            var htmlTags = formatCharsToHTML[result[1]];
            offset += (htmlTags[0].length + htmlTags[1].length) - (result[1].length*2);
        }
    }
}
const formatCharsToHTML = {
    "**": ["<b>", "</b>"]
}

window.onkeydown = function(e) {
    if (e.ctrlKey && e.key in shortcutKeys) {
        e.preventDefault();
        shortcutKeys[e.key]();
    }
}

function replaceRange(str, start, end, replace) {
    return str.substring(0, start) + replace + str.substring(end);
}
function genReplaceString(formatChars, substr) {
    var tags = formatCharsToHTML[formatChars];
    return tags[0] + substr + tags[1];
}