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

  modalDiv.find('#devToggleLink').click(() => {
    modalDiv.find('#devToggle').hide();
    modalDiv.find('#console').show();
  });

  runner($('#consoleInput'), $('#console ul.console-output'));
}

function runner(source: JQuery<HTMLElement>, destination: JQuery<HTMLElement>) {
  let appendOutput = (prefix: string, output: string) =>
    $('<li>', {text: `${prefix} ${output}`}).appendTo(destination);

  let run = (cmd: string) => {
    appendOutput('>', cmd);
    source.val('');
    try {
      appendOutput('<-', eval(cmd));
    } catch (err) {
      appendOutput('<!', err);
    }
    destination.scrollTop(destination.prop('scrollHeight'));
  }

  source.keypress(e => {
    if (e.which == 13) run(<string>source.val());
  })
}
