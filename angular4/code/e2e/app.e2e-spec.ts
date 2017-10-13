import { VotoUiPage } from './app.po';

describe('voto-ui App', () => {
  let page: VotoUiPage;

  beforeEach(() => {
    page = new VotoUiPage();
  });

  it('should display welcome message', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('Welcome to app!!');
  });
});
