updateSettings();
setInterval(updateSettings, 10000);

function reloadTables() {
    TeamsTable.ajax.reload(null, false);
    CategoriesTable.ajax.reload(null, false);
    ParticipantsTable.ajax.reload(null, false);
}

$(document).ready(function(){
    $(".form-control").after("<span></span>");
    $('.mdl-card__menu-button--add').mouseup(async function () {
        let button = $(this);
        let editInfo = button.data('content');
        button.toggle();
        switch (editInfo) {
            case "category": {
                CategoriesTable.row.add({id: "000", name: "Новая категория", negative_marks: "false"}).draw();
                $('.mdl-card__body-table-row__field#mdl-card__body-table-categories--name-000').dblclick();
                break;
            }
            case "team": {
                TeamsTable.row.add({id: "000", name: "Новая команда", leader: {login: ""}, participants: []}).draw();
                $('.mdl-card__body-table-row__field#mdl-card__body-table-teams--name-000').dblclick();
                break;
            }
            case "participant": {
                ParticipantsTable.row.add({
                    login: "000",
                    name: "Имя",
                    surname: "Фамилия",
                    middlename: "Отчество",
                    sex: 0,
                    team: 0
                }).draw();
                $('.mdl-card__body-table-row__field#mdl-card__body-table-participants--name-000').dblclick();
                break;
            }
        }
    });
    $('.mdl-card__menu-button--download').mousedown(function () {
        let button = $(this);
        let info = button.data('content');
        switch (info) {
            case "participants": {
                window.location.href = "https://api.forcamp.ga/orgset.participants.password.get?token="+Token;
                break;
            }
            case "employees": {
                window.location.href = "https://api.forcamp.ga/orgset.employees.password.get?token="+Token;
                break;
            }
        }
    });
});

function onTableDraw() {
    componentHandler.upgradeDom();
    $('.mdl-card__body-table-row__field:not(.mdl-card__body-table-row__field--noteditable)').unbind('dblclick').dblclick(function () {
        let textField = $(this);
        let baseText = textField.text();
        let editInfo = textField.data('content').split('-');
        textField.addClass('mdl-card__body-table-row__field--editing');
        textField.attr('contenteditable', 'true');
        textField.focus();
        textField.off('focusout keydown').on('focusout keydown',function (e) {
            if(e.keyCode === 13 || e.type === "focusout") {
                if (textField.attr('contenteditable') === "false") {
                    return
                }
                textField.removeClass('mdl-card__body-table-row__field--editing');
                textField.attr('contenteditable', 'false');
                let text = textField.text();
                if (text !== baseText && editInfo[1] !== "000") {
                    switch (editInfo[0]) {
                        case "category": {
                            let name = $('#mdl-card__body-table-categories--name-' + editInfo[1]).text();
                            let negative_marks = $('#mdl-card__body-table-categories--negative_marks-' + editInfo[1]).prop('checked');
                            editCategory(name, negative_marks, editInfo[1]);
                            break;
                        }
                        case "team": {
                            let name = $('#mdl-card__body-table-teams--name-'+editInfo[1]).text();
                            editTeam(name, editInfo[1]);
                            break;
                        }
                        case "participant": {
                            let name = $('#mdl-card__body-table-participants--name-'+editInfo[1]).text();
                            let surname = $('#mdl-card__body-table-participants--surname-'+editInfo[1]).text();
                            let middlename = $('#mdl-card__body-table-participants--middlename-'+editInfo[1]).text();
                            let sex = $('#mdl-card__body-table-participants--sex-'+editInfo[1]).data('content').split('-')[3];
                            let team = $('#mdl-card__body-table-participants--team-'+editInfo[1]).data('content').split('-')[3];
                            editParticipant(name, surname, middlename, sex, team, editInfo[1]);
                            break;
                        }
                    }
                }
            }
        });
    });
    $('.mdl-card__body-table-row_switch').unbind('mousedown').mousedown( function () {
        let toggle = $(this);
        let editInfo = toggle.attr('for').split('--')[1].split('-');
        let editObject = toggle.attr('for').split('--')[0].split('-')[3];
        if (editInfo[1] !== "000") {
            switch (editObject) {
                case "categories": {
                    let name = $('#mdl-card__body-table-categories--name-' + editInfo[1]).text();
                    let negative_marks = $('#mdl-card__body-table-categories--negative_marks-' + editInfo[1]).prop('checked').toString();
                    editCategory(name, negative_marks, editInfo[1]);
                    break;
                }
            }
        }
    });
    $('.mdl-card__body-table-row__dropdown').unbind('dblclick').dblclick( function () {
        let dropdownWrapper = $(this);
        dropdownWrapper.addClass('mdl-card__body-table-row__dropdown--editing');
        dropdownWrapper.children('ul').children('li').unbind('mousedown').mousedown(function () {
            dropdownWrapper.removeClass('mdl-card__body-table-row__dropdown--editing');
            let field = $(this);
            let editInfo = field.data('content').split('-');
            dropdownWrapper.children('.mdl-card__body-table-row__dropdown-ttl').text(field.text());
            dropdownWrapper.data('content', field.data('content'));
            switch (editInfo[0]) {
                case "participant": {
                    let name = $('#mdl-card__body-table-participants--name-'+editInfo[1]).text();
                    let surname = $('#mdl-card__body-table-participants--surname-'+editInfo[1]).text();
                    let middlename = $('#mdl-card__body-table-participants--middlename-'+editInfo[1]).text();
                    let sex = $('#mdl-card__body-table-participants--sex-'+editInfo[1]).data('content').split('-')[3];
                    let team = $('#mdl-card__body-table-participants--team-'+editInfo[1]).data('content').split('-')[3];
                    editParticipant(name, surname, middlename, sex, team, editInfo[1]);
                    break;
                }
            }
        });
    });
    $('.mdl-card__body-table-row_actions--create').unbind('mousedown').mousedown(async function () {
       let button = $(this);
       let creationInfo = button.data('content').split('-');
       switch (creationInfo[0]) {
           case "team": {
               let name = $('#mdl-card__body-table-teams--name-'+creationInfo[1]).text();
               let id = await addTeam(name);
               if (id === -1) {
                   TeamsTable.row(button.parents('tr')).remove().draw();
               } else {
                   TeamsTable.row(button.parents('tr')).remove();
                   TeamsTable.row.add({id: id, name: name, leader: {login: ""}, participants: []}).draw();
               }
               break;
           }
           case "category": {
               let name = $('#mdl-card__body-table-categories--name-' + creationInfo[1]).text();
               let negative_marks = $('#mdl-card__body-table-categories--negative_marks-' + creationInfo[1]).prop('checked').toString();
               let id = await addCategory(name, negative_marks);
               if (id === -1) {
                   CategoriesTable.row(button.parents('tr')).remove().draw();
               } else {
                   CategoriesTable.row(button.parents('tr')).remove();
                   CategoriesTable.row.add({id: id, name: name, negative_marks: negative_marks}).draw();
               }
               break;
           }
           case "participant": {
               let name = $('#mdl-card__body-table-participants--name-'+creationInfo[1]).text();
               let surname = $('#mdl-card__body-table-participants--surname-'+creationInfo[1]).text();
               let middlename = $('#mdl-card__body-table-participants--middlename-'+creationInfo[1]).text();
               let sex = $('#mdl-card__body-table-participants--sex-'+creationInfo[1]).data('content').split('-')[3];
               let team = $('#mdl-card__body-table-participants--team-'+creationInfo[1]).data('content').split('-')[3];
               let login = addParticipant(name, surname, middlename, sex, team);
               if (login === -1) {
                   ParticipantsTable.row(button.parents('tr')).remove().draw();
               } else {
                   ParticipantsTable.row(button.parents('tr')).remove();
                   ParticipantsTable.row.add({
                       login: login,
                       name: name,
                       surname: surname,
                       middlename: middlename,
                       sex: sex,
                       team: team
                   }).draw();
               }
               break;
           }
       }
        $('.mdl-card__menu-button--add[data-content='+creationInfo[0]+']').toggle();
    });
    $('.mdl-card__body-table-row_actions--decline').unbind('mousedown').mousedown(async function () {
        let button = $(this);
        let creationInfo = button.data('content').split('-');
        switch (creationInfo[0]) {
            case "team": {
                TeamsTable.row(button.parents('tr')).remove().draw();
                break;
            }
            case "category": {
                CategoriesTable.row(button.parents('tr')).remove().draw();
                break;
            }
            case "participant": {
                ParticipantsTable.row(button.parents('tr')).remove().draw();
                break;
            }
        }
        $('.mdl-card__menu-button--add[data-content='+creationInfo[0]+']').toggle();
    });
    $('.mdl-card__body-table-row_actions--delete').unbind('mousedown').mousedown( function () {
        let button = $(this);
        let editInfo = button.data('content').split('-');
        switch (editInfo[0]){
            case "category": {
                deleteCategory(editInfo[1], button);
                break;
            }
            case "team": {
                deleteTeam(editInfo[1], button);
                break;
            }
            case "participant": {
                deleteParticipant(editInfo[1], button);
                break;
            }
        }
    });
    $('.mdl-card__body-table-row_actions--reset_password').unbind('mousedown').mousedown( function () {
        let button = $(this);
        let editInfo = button.data('content').split('-');
        switch (editInfo[0]){
            case "employee": {
                deleteCategory(editInfo[1]);
                break;
            }
            case "participant": {
                resetParticipantPassword(editInfo[1]);
                break;
            }
        }
    });
    $('.mdl-card__body-table-row_actions--profile').unbind('mousedown').mousedown( function () {
       let button = $(this);
       let login = button.data('content').split('-')[1];
       window.location.href = "https://forcamp.ga/profile?login=" + login;
    });
    ActivateWavesEffect();
}

// Teams

function addTeam(name) {
    return new Promise(resolve => {
        $.post(__AddTeamLink, { token: Token, name: name }, function (resp) {
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
                reloadTables();
                resolve(resp.message.id);
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
                resolve(-1);
            }
        });
    });
}

function editTeam(name, id) {
    $.post(__EditTeamLink, { token: Token, name: name, id: id}, function (resp) {
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

function deleteTeam(id, button) {
    $.post(__DeleteTeamLink, { token: Token, id: id }, function (resp) {
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            TeamsTable.row(button.parents('tr')).remove().draw();
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

let TeamsTable = $('#mdl-card__body-table-teams').DataTable({
    "ajax": {
        "url": __GetTeamsLink,
        "type": "GET",
        "data": {
            "token": Token,
        },
        "dataSrc": function (data) {
            Teams = data.message.teams;
            return data.message.teams;
        },
    },
    columnDefs: [
        {
            targets: 0,
            name: "name",
            className: 'mdl-data-table__cell--non-numeric',
            data: "name",
            render: function ( name, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-teams--name-'+row.id+'"' +
                    ' data-content="team-'+row.id+'-name">'+
                    name[0].toUpperCase()+name.substring(1)+'</div>';
            },
        },
        {
            targets: 1,
            name: "leader",
            className: 'mdl-data-table__cell--non-numeric',
            data: "leader",
            searchable: false,
            orderable: false,
            render: function ( leader, type, row, meta ) {
                if (leader.login.length > 0) {
                    return '<a class="mdl-card__body-table-row__field mdl-card__body-table-row__field--noteditable mdl-card__body-table-row__field--capitalize" ' +
                        'id="mdl-card__body-table-teams--leader-' + row.id + '"' +
                        ' data-content="team-' + row.id + '-leader" href="https://forcamp.ga/profile?login=' + leader.login + '">' +
                        leader.surname + ' ' + leader.name + ' ' + leader.middlename + '</a>';
                } else {
                    return '<div class="mdl-card__body-table-row__field mdl-card__body-table-row__field--noteditable mdl-card__body-table-row__field--capitalize" ' +
                        'id="mdl-card__body-table-teams--leader-' + row.id + '"' +
                        ' data-content="team-' + row.id + '-leader">отсутствует</div>';
                }
            },
        },
        {
            targets: 2,
            name: "participants",
            className: '',
            data: "participants",
            searchable: false,
            orderable: true,
            render: function ( participants, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field mdl-card__body-table-row__field--noteditable mdl-card__body-table-row__field--capitalize" ' +
                    'id="mdl-card__body-table-teams--participants-'+row.id+'"' +
                    ' data-content="team-'+row.id+'-participants">'+participants.length+'</div>';
            }
        },
        {
            targets: 3,
            name: "actions",
            className: 'mdl-card__body-table-row_actions',
            data: "id",
            searchable: false,
            orderable: false,
            render: function ( id, type, row, meta ) {
                if (row.id === "000") {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--create"' +
                        ' data-content="team-' + id + '"> ' +
                        '<i class="material-icons">done</i></button>' +
                        '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--decline"' +
                        ' data-content="team-' + id + '"> ' +
                        '<i class="material-icons">close</i></button>';
                } else {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--delete"' +
                        ' data-content="team-' + id + '"> ' +
                        '<i class="material-icons">delete_forever</i></button>';
                }
            }
        },
    ],
    language: {
        "processing": "Подождите...",
        "search": "Поиск:",
        "lengthMenu": "Показать _MENU_ записей",
        "info": "",
        "infoEmpty": "",
        "infoFiltered": "",
        "infoPostFix": "",
        "loadingRecords": "Загрузка команд...",
        "zeroRecords": "Команды отсутствуют.",
        "emptyTable": "Команды отсутствуют",
        "paginate": {
            "first": "Первая",
            "previous": "Пред.",
            "next": "След.",
            "last": "Последняя"
        },
        "aria": {
            "sortAscending": ": отсортировать по возрастанию",
            "sortDescending": ": отсортировать по убыванию"
        }
    },
    "drawCallback": function () {
        $('#mdl-card__body-table-teams').css('width', '100%');
        onTableDraw();
    },
    responsive: true,
});

// Settings

async function updateSettings() {
    await GetOrganizationSettings();
    $('.mdl-card__body-row-text[data-content=participant]').text(OrgSettings.participant);
    $('.mdl-card__body-row-text[data-content=period]').text(OrgSettings.period);
    $('.mdl-card__body-row-text[data-content=organization]').text(OrgSettings.organization);
    $('.mdl-card__body-row-text[data-content=team]').text(OrgSettings.team);
    $('#fc-orgset__main-self_marks').prop('checked', OrgSettings.self_marks);
}

function editOrganizationSettings() {
    $.post(__SetOrgSettingValueLink, { token: Token,
        participant: OrgSettings.participant,
        team: OrgSettings.team,
        organization: OrgSettings.organization,
        period: OrgSettings.period,
        self_marks: !OrgSettings.self_marks
    }, function (resp) {
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены!", time: 2})
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

$('.mdl-card__body-row-switch').mousedown(function () {
    let toggle = $(this);
    let content = toggle.data('content');
    OrgSettings[content] = $('#fc-orgset__main-'+content).prop('checked');
    editOrganizationSettings();
});

function submitSettingEdit(textField, editButton, acceptButton, declineButton, content, baseText) {
    textField.unbind('keydown');
    acceptButton.unbind('click');
    declineButton.unbind('click');
    textField.attr('contenteditable', 'false');
    textField.removeClass('mdl-card__body-row-text--editing');
    editButton.removeClass('mdl-card__body-row-button--off');
    acceptButton.addClass('mdl-card__body-row-button--off');
    declineButton.addClass('mdl-card__body-row-button--off');
    OrgSettings[content] = textField.text();
    if (OrgSettings[content] !== baseText) {
        editOrganizationSettings();
    }
}

function declineSettingEdit(textField, editButton, acceptButton, declineButton, baseText) {
    textField.unbind('keydown');
    acceptButton.unbind('click');
    declineButton.unbind('click');
    textField.attr('contenteditable', 'false');
    textField.removeClass('mdl-card__body-row-text--editing');
    editButton.removeClass('mdl-card__body-row-button--off');
    acceptButton.addClass('mdl-card__body-row-button--off');
    declineButton.addClass('mdl-card__body-row-button--off');
    textField.text(baseText);
}

$('.mdl-card__body-row-button--edit').click(function () {
    let editButton = $(this);
    let content = editButton.data('content');
    let acceptButton = $('.mdl-card__body-row-button--accept[data-content='+content+']');
    let declineButton = $('.mdl-card__body-row-button--decline[data-content='+content+']');
    let textField = $('.mdl-card__body-row-text[data-content='+content+']');
    let baseText = textField.text();
    textField.keydown(function (e) {
        if((e.keyCode) === 13) {
            submitSettingEdit(textField, editButton, acceptButton, declineButton, content, baseText);
        }
    });
    acceptButton.click(function () {
        submitSettingEdit(textField, editButton, acceptButton, declineButton, content, baseText);
    });
    declineButton.click(function () {
        declineSettingEdit(textField, editButton, acceptButton, declineButton, baseText);
    });
    textField.addClass('mdl-card__body-row-text--editing');
    textField.attr('contenteditable', 'true');
    textField.focus();
    editButton.addClass('mdl-card__body-row-button--off');
    acceptButton.removeClass('mdl-card__body-row-button--off');
    declineButton.removeClass('mdl-card__body-row-button--off');
});

// Categories

function addCategory(name, negative_marks) {
    return new Promise(resolve => {
        $.post(__AddCategoryLink, { token: Token, name: name, negative_marks: negative_marks }, function (resp) {
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
                reloadTables();
                resolve(resp.message.id);
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
                resolve(-1);
            }
        });
    });
}

function editCategory(name, negative_marks, id) {
    $.post(__EditCategoryLink, { token: Token, name: name, negative_marks: !negative_marks, id: id}, function (resp) {
       if(resp.code === 200) {
           notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
           reloadTables();
       } else {
           notie.alert({type: 3, text: resp.message.ru, time: 2});
       }
    });
}

function deleteCategory(id, button) {
    $.post(__DeleteCategoryLink, { token: Token, id: id }, function (resp) {
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            CategoriesTable.row(button.parents('td')).remove().draw();
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

let CategoriesTable = $('#mdl-card__body-table-categories').DataTable({
    "ajax": {
        "url": __GetCategoriesLink,
        "type": "GET",
        "data": {
            "token": Token,
        },
        "dataSrc": function (data) {
            Categories = data.message.categories;
            return data.message.categories;
        },
    },
    columnDefs: [
        {
            targets: 0,
            name: "name",
            className: 'mdl-data-table__cell--non-numeric',
            data: "name",
            render: function ( name, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-categories--name-'+row.id+'"' +
                    ' data-content="category-'+row.id+'-name">'+
                    name[0].toUpperCase()+name.substring(1)+'</div>';
            },
        },
        {
            targets: 1,
            name: "negative_marks",
            className: 'mdl-data-table__cell--non-numeric',
            data: "negative_marks",
            searchable: false,
            orderable: false,
            render: function ( negative_marks, type, row, meta ) {
                if(negative_marks === "true") {
                    return '<label class="mdl-switch mdl-js-switch mdl-js-ripple-effect mdl-card__body-table-row_switch"' +
                        'for="mdl-card__body-table-categories--negative_marks-' + row.id + '">' +
                        '<input type="checkbox" id="mdl-card__body-table-categories--negative_marks-' + row.id + '" class="mdl-switch__input" checked> ' +
                        '<span class="mdl-switch__label"></span></label>';
                } else {
                    return '<label class="mdl-switch mdl-js-switch mdl-js-ripple-effect mdl-card__body-table-row_switch"' +
                        'for="mdl-card__body-table-categories--negative_marks-' + row.id + '">' +
                        '<input type="checkbox" id="mdl-card__body-table-categories--negative_marks-' + row.id + '" class="mdl-switch__input"> ' +
                        '<span class="mdl-switch__label"></span></label>';
                }
            },
        },
        {
            targets: 2,
            name: "actions",
            className: 'mdl-card__body-table-row_actions',
            data: "id",
            searchable: false,
            orderable: false,
            render: function ( id, type, row, meta ) {
                if (id === "000") {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--create"' +
                        ' data-content="category-' + id + '"> ' +
                        '<i class="material-icons">done</i></button>' +
                        '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--decline"' +
                        ' data-content="category-' + id + '"> ' +
                        '<i class="material-icons">close</i></button>';
                } else {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--delete"' +
                        ' data-content="category-' + id + '"> ' +
                        '<i class="material-icons">delete_forever</i></button>';
                }
            }
        },
    ],
    language: {
        "processing": "Подождите...",
        "search": "Поиск:",
        "lengthMenu": "Показать _MENU_ записей",
        "info": "",
        "infoEmpty": "",
        "infoFiltered": "",
        "infoPostFix": "",
        "loadingRecords": "Загрузка категорий...",
        "zeroRecords": "Категории отсутствуют.",
        "emptyTable": "Категории отсутствуют",
        "paginate": {
            "first": "Первая",
            "previous": "Пред.",
            "next": "След.",
            "last": "Последняя"
        },
        "aria": {
            "sortAscending": ": отсортировать по возрастанию",
            "sortDescending": ": отсортировать по убыванию"
        }
    },
    "drawCallback": function () {
        $('#mdl-card__body-table-categories').css('width', '100%');
        onTableDraw();
    },
    responsive: true,
});

// Participants

function addParticipant(name, surname, middlename, sex, team) {
    return new Promise(resolve => {
        $.post(__AddParticipantLink, { token: Token,
            name: name,
            surname: surname,
            middlename: middlename,
            sex: sex,
            team: team}, function (resp) {
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
                reloadTables();
                resolve(resp.message.id);
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
                resolve(-1);
            }
        });
    });
}

function editParticipant(name, surname, middlename, sex, team, login) {
    $.post(__EditParticipantLink, { token: Token,
        name: name,
        surname: surname,
        middlename: middlename,
        sex: sex,
        team: team,
        login: login}, function (resp) {
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

function deleteParticipant(login, button) {
    $.post(__DeleteParticipantLink, { token: Token, login: login }, function (resp) {
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            ParticipantsTable.row(button.parents('td')).remove().draw();
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

function resetParticipantPassword(login) {
    $.post(__ResetParticipantPasswordLink, { token: Token, login: login }, function (resp) {
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Новый пароль: " + resp.message.password, time: 0});
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

let ParticipantsTable = $('#mdl-card__body-table-participants').DataTable({
    "ajax": {
        "url": __GetParticipantsLink,
        "type": "GET",
        "data": {
            "token": Token,
        },
        "dataSrc": function (data) {
            return data.message.participants;
        },
    },
    columnDefs: [
        {
            targets: 0,
            name: "surname",
            className: 'mdl-data-table__cell--non-numeric',
            data: "surname",
            render: function ( surname, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-participants--surname-'+row.login+'"' +
                    ' data-content="participant-'+row.login+'-surname">'+
                    surname[0].toUpperCase()+surname.substring(1)+'</div>';
            },
        },
        {
            targets: 1,
            name: "name",
            className: 'mdl-data-table__cell--non-numeric',
            data: "name",
            render: function ( name, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-participants--name-'+row.login+'"' +
                    ' data-content="participant-'+row.login+'-name">'+
                    name[0].toUpperCase()+name.substring(1)+'</div>';
            },
        },
        {
            targets: 2,
            name: "middlename",
            className: 'mdl-data-table__cell--non-numeric',
            data: "middlename",
            render: function ( middlename, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-participants--middlename-'+row.login+'"' +
                    ' data-content="participant-'+row.login+'-middlename">'+
                    middlename[0].toUpperCase()+middlename.substring(1)+'</div>';
            },
        },
        {
            targets: 3,
            name: "sex",
            className: 'mdl-data-table__cell--non-numeric',
            data: "sex",
            searchable: false,
            orderable: false,
            render: function ( sex, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__dropdown mdl-card__body-table-row__field--capitalize" ' +
                    'id="mdl-card__body-table-participants--sex-'+row.login+'"' +
                    ' data-content="participant-' + row.login + '-sex-' + sex + '">' +
                    '<div class="mdl-card__body-table-row__dropdown-ttl">'+GetSexByID(sex)+'</div><ul>'+
                    '<li class="mdl-card__body-table-row__dropdown-field" data-content="participant-' + row.login + '-sex-0"><span class="wave-effect">мужской</span></li>'+
                    '<li class="mdl-card__body-table-row__dropdown-field" data-content="participant-' + row.login + '-sex-1"><span class="wave-effect">женский</span></li>'+
                    '</ul></label></div>';
            },
        },
        {
            targets: 4,
            name: "team",
            className: 'mdl-data-table__cell--non-numeric',
            data: "team",
            searchable: true,
            orderable: true,
            render: function ( team_id, type, row, meta ) {
                let dropdown = '<div class="mdl-card__body-table-row__dropdown mdl-card__body-table-row__field--capitalize" ' +
                'id="mdl-card__body-table-participants--team-'+row.login+'"' +
                ' data-content="participant-' + row.login + '-team-' + team_id + '">' +
                '<div class="mdl-card__body-table-row__dropdown-ttl">' + GetTeamNameByID(team_id) +'</div><ul>';
                Teams.forEach(function (team) {
                    dropdown += '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="participant-' + row.login + '-team-' + team.id + '">' +
                        '<span class="wave-effect">' + team.name + '</span></li>';
                });
                return dropdown + '</ul></label></div>';
            }
        },
        {
            targets: 5,
            name: "actions",
            className: 'mdl-card__body-table-row_actions',
            data: "login",
            searchable: false,
            orderable: false,
            render: function ( login, type, row, meta ) {
                if (login === "000") {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--create"' +
                        ' data-content="participant-' + login + '"> ' +
                        '<i class="material-icons">done</i></button>' +
                        '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--decline"' +
                        ' data-content="participant-' + login + '"> ' +
                        '<i class="material-icons">close</i></button>';
                } else {
                    return '<button class="mdl-button mdl-button--primary mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--profile"' +
                        ' data-content="participant-' + login + '"> ' +
                        '<i class="material-icons">person</i></button>' +
                        '<button class="mdl-button mdl-button--primary mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--reset_password"' +
                        ' data-content="participant-' + login + '"> ' +
                        '<i class="material-icons">refresh</i></button>' +
                        '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--delete"' +
                        ' data-content="participant-' + login + '"> ' +
                        '<i class="material-icons">delete_forever</i></button>';
                }
            }
        },
    ],
    language: {
        "processing": "Подождите...",
        "search": "Поиск:",
        "lengthMenu": "Показать _MENU_ записей",
        "info": "",
        "infoEmpty": "",
        "infoFiltered": "",
        "infoPostFix": "",
        "loadingRecords": "Загрузка участников...",
        "zeroRecords": "Участники отсутствуют.",
        "emptyTable": "Участники отсутствуют",
        "paginate": {
            "first": "Первая",
            "previous": "Пред.",
            "next": "След.",
            "last": "Последняя"
        },
        "aria": {
            "sortAscending": ": отсортировать по возрастанию",
            "sortDescending": ": отсортировать по убыванию"
        }
    },
    "drawCallback": function () {
        $('#mdl-card__body-table-participants').css('width', '100%');
        onTableDraw();
    },
    responsive: true,
});
