{{ define "aside" }}
<form id="search" class="form-inline" role="search" action="/search" method="get">
    <div class="form-group">
    <label for="Search" class="sr-only">Search</label>
        <input type="text" id="Search" name="search" class="form-control" placeholder="Search">
    </div>
    <button type="submit" class="btn btn-default">
        <i class="icon-search"></i>
    </button>
</form>
<div class="panel panel-info">
    <div class="panel-heading">
        <h3 class="panel-title">
            <span><i class="icon-list"></i>  最近更新</span>
        </h3>
    </div>
    <div class="panel-body">
        for item in aside_title 
        <li>
            <a href="/topic/ item.slug ">
                <p> item.title </p>
            </a>
        </li>
         end 
    </div>
</div>
<div class="panel panel-danger">
    <div class="panel-heading">
    <h3 class="panel-title">
        <span><i class=" icon-tags"></i>  标签</span>
    </h3>
    </div>
    <div id="tag_a" class="panel-body list-group">
         for tag in getAllTags 
        <a href="/tag/ tag.tag_type "> tag.tag_type  [ tag["COUNT(id)"] ]</a>
         end 
    </div>
</div>
{{ end }}