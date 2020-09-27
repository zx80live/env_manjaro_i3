" [main]
set vb t_vb=
set expandtab
set tabstop=2
set softtabstop=2
set shiftwidth=2
set autoindent
set wildmenu
set mouse=a
set encoding=utf8
set ffs=unix,dos,mac " symbol of the next line
set nowrap

" [appearance]
set laststatus=2
set number
syntax on
set cursorline
set cursorcolumn
" hi CursorLine  cterm=NONE ctermbg=0F171F ctermfg=white guibg=darkred guifg=white
" hi CursorLine  cterm=NONE ctermbg=808080 ctermfg=white guibg=darkred guifg=white
colorscheme darcula
let g:airline_powerline_fonts = 1
set t_Co=256


" [search]
set ignorecase
set smartcase
set hlsearch
set incsearch

" [clipboard]
set clipboard=unnamed

" [plugins]
"    Repository         https://vimawesome.com/
"    Plugin manager     https://github.com/junegunn/vim-plug
if empty(glob('~/.vim/autoload/plug.vim'))
  silent !curl -fLo ~/.vim/autoload/plug.vim --create-dirs
    \ https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
  autocmd VimEnter * PlugInstall --sync | source $MYVIMRC
endif

call plug#begin('~/.vim/bundle')

Plug 'jiangmiao/auto-pairs'
Plug 'valloric/youcompleteme'
Plug 'tpope/vim-fugitive'
Plug 'airblade/vim-gitgutter'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'lilydjwg/colorizer'

call plug#end()


" [keys]
" map <C-n> :NERDTreeToggle<CR>
nmap <F3> i<C-R>=strftime("%H:%M - ")<CR><Esc>
imap <F3> <C-R>=strftime("%H:%M - ")<CR>
:inoremap <C-v> <ESC>"+pa
:vnoremap <C-c> "+y
