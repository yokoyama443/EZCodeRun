class MenuComponent extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
    }

    async connectedCallback() {
        await this.render();
        this.addEventListeners();
    }

    async render() {
        const styleSheet = await this.loadStylesheet();
        const html = this.getHTML();
        
        this.shadowRoot.adoptedStyleSheets = [styleSheet];
        this.shadowRoot.innerHTML = html;
    }

    async loadStylesheet() {
        const response = await fetch('css/menu.css');
        const cssText = await response.text();
        const styleSheet = new CSSStyleSheet();
        await styleSheet.replace(cssText);
        return styleSheet;
    }

    getHTML() {
        return `
            <button class="menu-button">
                <span></span>
                <span></span>
                <span></span>
            </button>
            <div class="popup-menu">
                <ul class="menu-items">
                    <li><a href="index.html">ホーム</a></li>
                    <li><a href="problems.html">問題一覧</a></li>
                    <li><a href="login.html">ログイン</a></li>
                    <li><a href="register.html">アカウント登録</a></li>
                </ul>
            </div>
        `;
    }

    addEventListeners() {
        const menuButton = this.shadowRoot.querySelector('.menu-button');
        const popupMenu = this.shadowRoot.querySelector('.popup-menu');

        menuButton.addEventListener('click', () => {
            menuButton.classList.toggle('active');
            popupMenu.classList.toggle('active');
        });

        popupMenu.addEventListener('click', (e) => {
            if (e.target.tagName === 'A') {
                menuButton.classList.remove('active');
                popupMenu.classList.remove('active');
            }
        });
    }
}

customElements.define('menu-component', MenuComponent);