export namespace main {
	
	export class Note {
	    id: string;
	    title: string;
	    content: string;
	    tags: string[];
	    path: string;
	    created: string;
	    modified: string;
	
	    static createFrom(source: any = {}) {
	        return new Note(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.content = source["content"];
	        this.tags = source["tags"];
	        this.path = source["path"];
	        this.created = source["created"];
	        this.modified = source["modified"];
	    }
	}

}

