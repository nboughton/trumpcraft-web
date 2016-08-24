{{ template "includes/header" }}
  <div class="container">
    <article>
      <p>What happens when you feed Donald Trump speeches and short stories by HP
        Lovecraft into a Markov chain generator? Well... This:
      </p>
      <label>Source</label>
	    <select id="trumpcraft-source">
		    <option value="trumpcraft">Trumpcraft</option>
        <option value="peter_rabbitcraft">Peter Rabbitcraft</option>
		    <option value="trump">Trump</option>
		    <option value="lovecraft">Lovecraft</option>
        <option value="peter_rabbit">Peter Rabbit</option>
        <option value="all">All</option>
	    </select>
	    <label>No. of words (max)</label>
      <select id="trumpcraft-words">
        <option value="100">100</option>
        <option value="250">250</option>
        <option value="500">500</option>
        <option value="750">750</option>
        <option value="1000">1000</option>
      </select>
	    <button id="trumpcraft-submit">Fhtagn!</button>
	    <div id="trumpcraft-results"></div>
    </article>
  </div>
{{ template "includes/footer" }}

