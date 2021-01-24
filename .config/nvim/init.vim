" [plugins] -------------------------------------------------------------
call plug#begin()

" themes
Plug 'blueshirts/darcula'
Plug 'joshdick/onedark.vim'
Plug 'arcticicestudio/nord-vim'
Plug 'christianchiarulli/nvcode-color-schemes.vim'
Plug 'jacoborus/tender.vim'
Plug 'nanotech/jellybeans.vim'
Plug 'mhartington/oceanic-next'
Plug 'morhetz/gruvbox'
" ui
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'ryanoasis/vim-devicons'

" navigation
Plug 'preservim/nerdtree'

" dev
Plug 'lilydjwg/colorizer'
Plug 'tpope/vim-fugitive'
Plug 'sheerun/vim-polyglot'
Plug 'derekwyatt/vim-scala'
Plug 'neoclide/coc.nvim', {'branch': 'release'}
Plug 'scalameta/coc-metals', {'do': 'yarn install --frozen-lockfile'}

call plug#end()

" [keys] ----------------------------------------------------------------
nnoremap <leader>n :NERDTreeFocus<CR>
nnoremap <C-n> :NERDTree<CR>
nnoremap <C-t> :NERDTreeToggle<CR>
nnoremap <C-f> :NERDTreeFind<CR>

" [appearance] ----------------------------------------------------------
set laststatus=2
set number
syntax on
set cursorline
set cursorcolumn
" hi CursorLine  cterm=NONE ctermbg=0F171F ctermfg=white guibg=darkred guifg=white
" hi CursorLine  cterm=NONE ctermbg=808080 ctermfg=white guibg=darkred guifg=white
let g:airline_powerline_fonts = 1
set t_Co=256
set encoding=UTF-8

if (has("termguicolors"))
 set termguicolors
endif

let g:oceanic_next_terminal_bold = 1
let g:oceanic_next_terminal_italic = 1
"let g:airline_theme='oceanicnext'
"let g:airline_theme='deus'
let g:airline_theme='onedark'
"let g:airline_theme = 'tender'
"let g:airline_theme='nord'

"colorscheme OceanicNext
colorscheme onedark
"colorscheme nord
"colorscheme gruvbox
"colorscheme tender 
"colorscheme jellybeans

" [main] ----------------------------------------------------------------
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

" [search] --------------------------------------------------------------
set ignorecase
set smartcase
set hlsearch
set incsearch

" [clipboard] -----------------------------------------------------------
set clipboard=unnamed

" [dev]
let g:scala_scaladoc_indent = 1
autocmd FileType json syntax match Comment +\/\/.\+$+

