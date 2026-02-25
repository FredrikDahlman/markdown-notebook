<script>
  import { onMount } from 'svelte';
  import { ListNotes, GetNote, SaveNote, CreateNote, DeleteNote, SearchNotes, GetAllTags, GetNotesByTag } from '../wailsjs/go/main/App.js';
  import { EditorView, basicSetup } from 'codemirror';
  import { EditorState } from '@codemirror/state';
  import { markdown, markdownLanguage } from '@codemirror/lang-markdown';
  import { languages } from '@codemirror/language-data';
  import { keymap } from '@codemirror/view';
  import { defaultKeymap } from '@codemirror/commands';
  import { HighlightStyle, syntaxHighlighting } from '@codemirror/language';
  import { tags as lezerTags } from '@lezer/highlight';
  import MarkdownIt from 'markdown-it';
  import { createHighlighter } from 'shiki';

  let notes = [];
  let currentNote = null;
  let editorView = null;
  let editorContainer;
  let previewContainer;
  let searchQuery = '';
  let tags = {};
  let selectedTag = null;
  let saveTimeout = null;
  let md = new MarkdownIt();
  let splitPosition = 50;
  let isDragging = false;

  let highlighter = null;

  function startDrag(e) {
    isDragging = true;
    document.addEventListener('mousemove', onDrag);
    document.addEventListener('mouseup', stopDrag);
  }

  function onDrag(e) {
    if (!isDragging) return;
    const main = document.querySelector('.main');
    if (!main) return;
    const rect = main.getBoundingClientRect();
    const newPosition = ((e.clientX - rect.left) / rect.width) * 100;
    splitPosition = Math.max(20, Math.min(80, newPosition));
  }

  function stopDrag() {
    isDragging = false;
    document.removeEventListener('mousemove', onDrag);
    document.removeEventListener('mouseup', stopDrag);
  }

  onMount(async () => {
    highlighter = await createHighlighter({
      themes: ['light-plus'],
      langs: ['javascript', 'typescript', 'go', 'python', 'rust', 'html', 'css', 'json', 'yaml', 'bash', 'markdown']
    });

    await loadNotes();
    await loadTags();
  });

  async function loadNotes() {
    let loadedNotes;
    if (selectedTag) {
      loadedNotes = await GetNotesByTag(selectedTag);
    } else if (searchQuery) {
      loadedNotes = await SearchNotes(searchQuery);
    } else {
      loadedNotes = await ListNotes();
    }
    notes = loadedNotes || [];
  }

  async function loadTags() {
    tags = await GetAllTags() || {};
  }

  async function openNote(note) {
    if (saveTimeout) {
      clearTimeout(saveTimeout);
      await saveCurrentNote();
    }

    currentNote = await GetNote(note.path);
    initEditor(currentNote.content);
    updatePreview();
  }

  function initEditor(content) {
    if (editorView) {
      editorView.destroy();
    }

    const markdownHighlight = HighlightStyle.define([
      { tag: lezerTags.heading, color: '#569cd6', fontWeight: 'bold' },
      { tag: lezerTags.heading1, color: '#569cd6', fontWeight: 'bold', fontSize: '1.5em' },
      { tag: lezerTags.heading2, color: '#569cd6', fontWeight: 'bold', fontSize: '1.3em' },
      { tag: lezerTags.heading3, color: '#569cd6', fontWeight: 'bold' },
      { tag: lezerTags.emphasis, fontStyle: 'italic', color: '#ce9178' },
      { tag: lezerTags.strong, fontWeight: 'bold', color: '#ce9178' },
      { tag: lezerTags.link, color: '#4ec9b0', textDecoration: 'underline' },
      { tag: lezerTags.url, color: '#4ec9b0' },
      { tag: lezerTags.monospace, color: '#ce9178', backgroundColor: '#3c3c3c' },
      { tag: lezerTags.quote, color: '#6a9955', fontStyle: 'italic' },
      { tag: lezerTags.list, color: '#dcdcaa' },
    ]);

    const state = EditorState.create({
      doc: content || '',
      extensions: [
        basicSetup,
        markdown({ base: markdownLanguage, codeLanguages: languages }),
        keymap.of(defaultKeymap),
        syntaxHighlighting(markdownHighlight),
        EditorView.updateListener.of((update) => {
          if (update.docChanged) {
            updatePreviewDebounced();
          }
        }),
        EditorView.theme({
          '&': { height: '100%', backgroundColor: '#1e1e1e' },
          '.cm-scroller': { overflow: 'auto' },
          '.cm-content': { fontFamily: 'Menlo, Monaco, monospace', fontSize: '14px', color: '#d4d4d4' },
          '.cm-gutters': { backgroundColor: '#252526', borderRight: '1px solid #3c3c3c' },
          '.cm-activeLineGutter': { backgroundColor: '#2a2a2a' },
          '.cm-activeLine': { backgroundColor: '#2a2a2a' },
          '.cm-cursor': { borderLeftColor: '#d4d4d4' },
          '.cm-selectionBackground': { backgroundColor: '#264f78' },
          '&.cm-focused .cm-selectionBackground': { backgroundColor: '#264f78' }
        })
      ]
    });

    editorView = new EditorView({
      state,
      parent: editorContainer
    });
  }

  function updatePreviewDebounced() {
    if (saveTimeout) {
      clearTimeout(saveTimeout);
    }
    saveTimeout = setTimeout(async () => {
      if (currentNote) {
        currentNote.content = editorView.state.doc.toString();
        await SaveNote(currentNote);
      }
    }, 500);

    updatePreview();
  }

  async function updatePreview() {
    if (!editorView || !previewContainer) return;

    const content = editorView.state.doc.toString();
    let html = md.render(content);

    // Process code blocks with Shiki
    const codeBlockRegex = /<pre><code class="language-(\w+)">([\s\S]*?)<\/code><\/pre>/g;
    html = html.replace(codeBlockRegex, (match, lang, code) => {
      try {
        const decoded = code
          .replace(/</g, '<')
          .replace(/>/g, '>')
          .replace(/&/g, '&')
          .replace(/"/g, '"')
          .replace(/'/g, "'");
        const highlighted = highlighter.codeToHtml(decoded, { lang: lang || 'text', theme: 'light-plus' });
        return highlighted;
      } catch (e) {
        return match;
      }
    });

    previewContainer.innerHTML = html;
  }

  let isSyncing = false;

  function syncEditorScroll(scrollInfo) {
    if (isSyncing || !previewContainer) return;
    isSyncing = true;
    const editorScrollInfo = scrollInfo;
    const previewScrollRatio = editorScrollInfo.top / (editorScrollInfo.height - editorScrollInfo.clientHeight);
    const previewMaxScroll = previewContainer.scrollHeight - previewContainer.clientHeight;
    previewContainer.scrollTop = previewMaxScroll * previewScrollRatio;
    setTimeout(() => isSyncing = false, 50);
  }

  function syncPreviewScroll() {
    if (isSyncing || !editorView || !previewContainer) return;
    isSyncing = true;
    const previewScrollRatio = previewContainer.scrollTop / (previewContainer.scrollHeight - previewContainer.clientHeight);
    const editorScrollInfo = editorView.scrollDOM;
    const editorMaxScroll = editorScrollInfo.scrollHeight - editorScrollInfo.clientHeight;
    editorView.scrollDOM.scrollTop = editorMaxScroll * previewScrollRatio;
    setTimeout(() => isSyncing = false, 50);
  }

  async function saveCurrentNote() {
    if (currentNote && editorView) {
      currentNote.content = editorView.state.doc.toString();
      await SaveNote(currentNote);
      await loadNotes();
      await loadTags();
    }
  }

  async function handleNewNote() {
    if (saveTimeout) {
      clearTimeout(saveTimeout);
      await saveCurrentNote();
    }

    currentNote = await CreateNote();
    await loadNotes();
    await loadTags();
    initEditor(currentNote.content);
    updatePreview();
  }

  let showDeleteConfirm = false;
  let isEditingTitle = false;
  let editedTitle = '';

  function startEditTitle() {
    if (!currentNote) return;
    editedTitle = currentNote.title;
    isEditingTitle = true;
  }

  async function saveTitle() {
    if (!currentNote || !editedTitle.trim()) {
      isEditingTitle = false;
      return;
    }
    currentNote.title = editedTitle.trim();
    await SaveNote(currentNote);
    await loadNotes();
    isEditingTitle = false;
  }

  function cancelEditTitle() {
    isEditingTitle = false;
    editedTitle = '';
  }

  function handleTitleKeydown(e) {
    if (e.key === 'Enter') {
      saveTitle();
    } else if (e.key === 'Escape') {
      cancelEditTitle();
    }
  }

  async function handleDeleteNote() {
    console.log('Delete clicked, currentNote:', currentNote);
    if (!currentNote) {
      alert('No note selected');
      return;
    }
    showDeleteConfirm = true;
  }

  async function confirmDelete() {
    showDeleteConfirm = false;
    try {
      console.log('Calling DeleteNote with path:', currentNote.path);
      await DeleteNote(currentNote.path);
      console.log('DeleteNote completed');
      currentNote = null;
      if (editorView) {
        editorView.destroy();
        editorView = null;
      }
      await loadNotes();
      await loadTags();
      alert('Note deleted!');
    } catch (err) {
      console.error('Delete failed:', err);
      alert('Failed to delete note: ' + err);
    }
  }

  function cancelDelete() {
    showDeleteConfirm = false;
  }

  async function handleSearch() {
    selectedTag = null;
    await loadNotes();
  }

  async function handleTagClick(tag) {
    selectedTag = tag;
    searchQuery = '';
    await loadNotes();
  }

  async function clearFilter() {
    selectedTag = null;
    searchQuery = '';
    await loadNotes();
  }

  function formatDate(dateStr) {
    if (!dateStr) return '';
    const date = new Date(dateStr);
    return date.toLocaleDateString();
  }
</script>

<div class="app">
  <div class="sidebar">
    <div class="sidebar-header">
      <input 
        type="text" 
        placeholder="Search notes..." 
        bind:value={searchQuery}
        on:input={handleSearch}
      />
      <button on:click={handleNewNote}>+ New</button>
    </div>

    {#if selectedTag || searchQuery}
      <div class="filter-info">
        {#if selectedTag}
          <span>Tag: {selectedTag}</span>
        {/if}
        {#if searchQuery}
          <span>Search: {searchQuery}</span>
        {/if}
        <button class="clear-btn" on:click={clearFilter}>Clear</button>
      </div>
    {/if}

    <div class="notes-list">
      {#each notes as note}
        <div 
          class="note-item" 
          class:active={currentNote && currentNote.path === note.path}
          on:click={() => openNote(note)}
        >
          <div class="note-title">{note.title || 'Untitled'}</div>
          <div class="note-meta">
            <span>{formatDate(note.modified)}</span>
            {#if note.tags && note.tags.length > 0}
              <span class="note-tags">{note.tags.slice(0, 2).join(', ')}</span>
            {/if}
          </div>
        </div>
      {/each}
    </div>

    <div class="tags-section">
      <div class="tags-header">Tags</div>
      <div class="tags-list">
        {#each Object.entries(tags) as [tag, count]}
          <button 
            class="tag" 
            class:selected={selectedTag === tag}
            on:click={() => handleTagClick(tag)}
          >
            {tag} ({count})
          </button>
        {/each}
      </div>
    </div>
  </div>

  <div class="main">
    <div class="editor-pane" style="width: {splitPosition}%">
      <div class="pane-header">
        {#if currentNote}
          {#if isEditingTitle}
            <input 
              class="title-input" 
              bind:value={editedTitle} 
              on:blur={saveTitle}
              on:keydown={handleTitleKeydown}
              autofocus
            />
          {:else}
            <span class="title-span" on:dblclick={startEditTitle}>{currentNote.title}</span>
          {/if}
          <button class="delete-btn" on:click={() => handleDeleteNote()}>Delete</button>
        {:else}
          <span>Select a note</span>
        {/if}
      </div>
      <div class="editor" bind:this={editorContainer}></div>
    </div>

    <div class="divider" on:mousedown={startDrag}></div>

    <div class="preview-pane" style="width: {100 - splitPosition}%">
      <div class="pane-header">Preview</div>
      <div class="preview" bind:this={previewContainer} on:scroll={syncPreviewScroll}></div>
    </div>
  </div>

  {#if showDeleteConfirm}
    <div class="modal-overlay">
      <div class="modal">
        <p>Are you sure you want to delete "{currentNote?.title}"?</p>
        <div class="modal-buttons">
          <button class="cancel-btn" on:click={cancelDelete}>Cancel</button>
          <button class="confirm-delete-btn" on:click={confirmDelete}>Delete</button>
        </div>
      </div>
    </div>
  {/if}
</div>

<style>
  :global(body) {
    margin: 0;
    padding: 0;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, sans-serif;
    overflow: hidden;
  }

  .app {
    display: flex;
    height: 100vh;
    width: 100vw;
  }

  .sidebar {
    width: 280px;
    min-width: 200px;
    background: #f5f5f5;
    border-right: 1px solid #ddd;
    display: flex;
    flex-direction: column;
  }

  .sidebar-header {
    padding: 12px;
    border-bottom: 1px solid #ddd;
    display: flex;
    gap: 8px;
  }

  .sidebar-header input {
    flex: 1;
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 4px;
  }

  .sidebar-header button {
    padding: 8px 12px;
    background: #007aff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }

  .filter-info {
    padding: 8px 12px;
    background: #e8f4ff;
    font-size: 12px;
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .clear-btn {
    padding: 2px 6px;
    font-size: 10px;
    background: #ddd;
    border: none;
    border-radius: 2px;
    cursor: pointer;
    margin-left: auto;
  }

  .notes-list {
    flex: 1;
    overflow-y: auto;
  }

  .note-item {
    padding: 12px;
    border-bottom: 1px solid #eee;
    cursor: pointer;
  }

  .note-item:hover {
    background: #eaeaea;
  }

  .note-item.active {
    background: #d4e5ff;
  }

  .note-title {
    font-weight: 600;
    margin-bottom: 4px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    color: #1a1a1a;
  }

  .note-meta {
    font-size: 11px;
    color: #666;
    display: flex;
    gap: 8px;
  }

  .note-tags {
    color: #007aff;
  }

  .tags-section {
    border-top: 1px solid #ddd;
    padding: 12px;
    max-height: 200px;
    overflow-y: auto;
  }

  .tags-header {
    font-weight: 500;
    margin-bottom: 8px;
    font-size: 12px;
    color: #666;
  }

  .tags-list {
    display: flex;
    flex-wrap: wrap;
    gap: 4px;
  }

  .tag {
    padding: 2px 8px;
    font-size: 11px;
    background: #eee;
    border: none;
    border-radius: 10px;
    cursor: pointer;
  }

  .tag.selected {
    background: #007aff;
    color: white;
  }

  .main {
    flex: 1;
    display: flex;
  }

  .editor-pane, .preview-pane {
    display: flex;
    flex-direction: column;
    min-width: 0;
  }

  .divider {
    width: 6px;
    background: #ddd;
    cursor: col-resize;
    flex-shrink: 0;
    transition: background 0.2s;
  }

  .divider:hover {
    background: #007aff;
  }

  .pane-header {
    padding: 8px 12px;
    background: #2d2d2d;
    border-bottom: 1px solid #444;
    font-weight: 600;
    font-size: 13px;
    color: #e0e0e0;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .title-span {
    cursor: pointer;
    padding: 2px 4px;
    border-radius: 3px;
  }

  .title-span:hover {
    background: #444;
  }

  .title-input {
    background: #444;
    color: white;
    border: 1px solid #666;
    border-radius: 3px;
    padding: 2px 6px;
    font-size: 13px;
    font-weight: 600;
    flex: 1;
    margin-right: 8px;
  }

  .delete-btn {
    padding: 4px 8px;
    background: #ff3b30;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 11px;
    z-index: 1000;
    position: relative;
  }

  .editor {
    flex: 1;
    overflow: hidden;
  }

  .editor :global(.cm-editor) {
    height: 100%;
  }

  .preview {
    flex: 1;
    padding: 20px;
    overflow-y: auto;
    background: #fafafa;
    color: #1a1a1a;
  }

  .preview :global(h1) {
    font-size: 2em;
    margin-bottom: 0.5em;
    border-bottom: 1px solid #eee;
    padding-bottom: 0.3em;
  }

  .preview :global(h2) {
    font-size: 1.5em;
    margin-top: 1em;
    margin-bottom: 0.5em;
  }

  .preview :global(p) {
    margin-bottom: 1em;
    line-height: 1.6;
  }

  .preview :global(code) {
    background: #f4f4f4;
    padding: 2px 4px;
    border-radius: 3px;
    font-size: 0.9em;
  }

  .preview :global(pre) {
    background: #f4f4f4;
    padding: 12px;
    border-radius: 6px;
    overflow-x: auto;
  }

  .preview :global(pre code) {
    background: none;
    padding: 0;
  }

  .preview :global(ul), .preview :global(ol) {
    margin-bottom: 1em;
    padding-left: 2em;
  }

  .preview :global(li) {
    margin-bottom: 0.25em;
  }

  .preview :global(blockquote) {
    border-left: 4px solid #ddd;
    margin: 1em 0;
    padding-left: 1em;
    color: #666;
  }

  .preview :global(a) {
    color: #007aff;
  }

  .preview :global(img) {
    max-width: 100%;
  }

  .preview :global(table) {
    border-collapse: collapse;
    width: 100%;
    margin-bottom: 1em;
  }

  .preview :global(th), .preview :global(td) {
    border: 1px solid #ddd;
    padding: 8px;
    text-align: left;
  }

  .preview :global(th) {
    background: #f5f5f5;
  }

  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 2000;
  }

  .modal {
    background: #333;
    color: white;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
    max-width: 400px;
  }

  .modal p {
    margin: 0 0 20px 0;
    font-size: 16px;
    color: white;
  }

  .modal-buttons {
    display: flex;
    gap: 10px;
    justify-content: flex-end;
  }

  .cancel-btn {
    padding: 8px 16px;
    background: #eee;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }

  .confirm-delete-btn {
    padding: 8px 16px;
    background: #ff3b30;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
</style>