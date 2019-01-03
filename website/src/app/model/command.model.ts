export class Command {
  id:           number;
  regex:        string;
  description:  string;
  response:     string;
  script:       boolean;
}

export function isValid(cmd: Command): boolean {
  if (cmd.regex.startsWith('*') || cmd.regex.startsWith('^*') || cmd.regex.startsWith('+')) {
    return false;
  }
  return true;
}
