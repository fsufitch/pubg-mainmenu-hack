let evilForm = require('./evil_form.html');
let postHack = require('./post_hack.html');

export function createPageStructure(apiHost: string, cbFinished: () => void) {
  let modalDiv = $('<div>', {id: "evilModal"});
  modalDiv.html(evilForm);

  modalDiv.find('#evilSubmit').click(() => {
    let username = modalDiv.find('#evilUsername').val();
    let password = modalDiv.find('#evilPassword').val();
    modalDiv.remove();
    renderPostHack(apiHost, cbFinished);
  });

  $('#content').append(modalDiv);
}

export function renderPostHack(apiHost: string, cbFinished: () => void) {
  let modalDiv = $('<div>', {id: "evilModal"});
  modalDiv.html(postHack)
  $('#content').append(modalDiv);

  console.log(modalDiv.find('#apiUrl'));
  console.log(apiHost);
  modalDiv.find('#apiUrl').html(apiHost);
  modalDiv.find('#continue').click(() => cbFinished());
}
