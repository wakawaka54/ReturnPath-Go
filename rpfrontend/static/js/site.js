// Write your Javascript code.

const limit = 20;
var offset = 0;
var totalCount = 0;
var currentPage = 0;
var totalPages = 0;
var scrollY = 0;

$(ajaxAddSentence_Setup());
$(ajaxShowSentence_Setup());

function formatPages()
{
    let pageElement = $('.pagination');
    let maxPageShown = 5;
    let pagesShownPerSection = 5;
    let i = 1;

    if(currentPage > Math.ceil(pagesShownPerSection / 2)) {
        maxPageShown = currentPage + Math.ceil(pagesShownPerSection / 2);
    }

    pageElement.empty();

    for(let p = maxPageShown - pagesShownPerSection; p < maxPageShown && p < totalPages; p++)
    {
        let li = '<li>';
        let a = '<a onclick="ajaxShowSentences(' + p + ')">' + (p + 1) + '</a>';
        if (currentPage == p) { li = '<li class="active">'; }
        if (p >= totalPages) { li = '<li class="disabled">'; }

        pageElement.append(li + a + '</li>');
    }
}

function ajaxAddSentence_Setup()
{
    $('#btnAddSentence').on('click', ajaxAddSentence);
}

function ajaxAddSentence(event)
{
    let serviceUrl = apiRoute;
    let form = $('#fAddSentence');

    $.ajax({
        type: 'POST',
        contentType: "application/json",
        data: toJson(form),
        url: serviceUrl,
        success: successFunction,
        error: errorFunction
    });

    event.preventDefault();

    function successFunction(data, status)
    {
        $('#sentenceValidation').empty();

        ajaxShowSentences();
    }

    function errorFunction(err)
    {
        console.log(err);

        $('#sentenceValidation').append(err.responseText);
    }
}

function ajaxShowSentence_Setup()
{
    ajaxShowSentences();
}

function ajaxShowSentences(page)
{
    if(page == undefined) {
        page = 0;
    }
    currentPage = page;
    offset = page * limit;

    let serviceUrl = apiRoute + "?limit=20&offset=" + offset;

    $.ajax({
        type: 'GET',
        url: serviceUrl,
        dataType: "json",
        success: successFunction,
        error: function(err) { console.log(err); }
    });

    function successFunction(data, status, xhr)
    {
        totalCount = xhr.getResponseHeader('X-Total-Count');
        totalPages = totalCount / limit;

        formatPages();

        let tbody = $('#tbSentences tbody');
        $(tbody).empty();

        if (data.length == 0)
        {
            $(tbody).append('<tr><td>There seems to be nothing here</td></tr>');
        }

        for (var i = 0; i != data.length; i++)
        {
            tags = "";
            tags = data[i].tags.join(", ")

            $(tbody).append('<tr><td>' + data[i].sentence + '</td><td style="white-space:pre-wrap;word-wrap:break-word;">' + tags + '</td><td><a class="btn btn-danger pull-right" onclick="ajaxDeleteSentence(' + "'" + data[i].id + "'" + ')">x</a></td></tr>');
        }

        ajaxStatistics();
    }
}

function ajaxDeleteSentence(id)
{
    let serviceUrl = apiRoute + "/" + id;

    scrollY = $('body').scrollTop()

    $.ajax({
        type: 'DELETE',
        url: serviceUrl,
        success: successFunction,
        error: function (err) { console.log(err); }
    });

    function successFunction(data, status) {
        ajaxShowSentences();
    }
}

function ajaxStatistics()
{
    let serviceUrl = apiRoute + "/statistics";

    $.ajax({
        type: 'GET',
        url: serviceUrl,
        dataType: "json",
        success: successFunction,
        error: function (err) { console.log(err); }
    });

    function successFunction(data, status) {
        chartReady(data);
        $('body').scrollTop(scrollY)
    }
}

function chartReady(data) {

    $('#barchart').remove();
    $('#barchart-container').empty();
    $('#barchart-container').append('<canvas id="barchart" width="200" height="200"></canvas>');

    var dataArray = [],
        dataLabels = [],
        dataValues = [],
        colors = [];

    for (var i = 0; i < data.length; i++) {
        dataArray.push([data[i].tag, data[i].count]);
    }

    dataArray.sort(function (a, b) {
        return b[1] - a[1];
    });

    dataArray.forEach(function (call, index) {
        let point = dataArray[index];
        dataLabels.push(point[0]);
        dataValues.push(point[1]);
        colors.push(randomColor());
    });

    var ctx = $('#barchart');
    var myChart = new Chart(ctx, {
        type: 'horizontalBar',
        data: {
            labels: dataLabels,
            datasets: [{
                label: 'Word Frequency',
                data: dataValues,
                backgroundColor: colors,
            }]
        },
        options: {
            scales: {
                xAxes: [{
                    ticks: {
                        beginAtZero: true
                    }

                }]
            },
            animation: {
                duration: 0
            }
        }
    });
}

function compareDataPoints(a, b) {
    if (a.point < b.point)
        return -1;
    if (a.point > b.point)
        return 1;
    return 0;
}

function randomColor() {
    var r = Math.ceil(Math.random() * 80);
    var g = Math.ceil(Math.random() * 80);
    var b = Math.ceil(Math.random() * 80);

    return 'rgba(' + r.toString() + ',' + g.toString() + ',' + b.toString() + ',.2)';
}

function toJson(form)
{
    let formObject = form.serializeArray();
    let formData = {};

    for(var i = 0; i < formObject.length; i++)
    {
        formData[formObject[i].name] = formObject[i].value;
    }

    return JSON.stringify(formData);
}
