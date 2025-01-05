document.addEventListener("DOMContentLoaded", () => {
    const expenseForm = document.getElementById("expense-form");
    const expensesTableBody = document.querySelector("#expenses-table tbody");
    const categorySummary = document.getElementById("category-summary");
    const chartContainer = document.getElementById("expense-chart");
  
    let expenses = JSON.parse(localStorage.getItem("expenses")) || [];
  
    const renderExpenses = () => {
      expensesTableBody.innerHTML = "";
      expenses.forEach((expense, index) => {
        const row = document.createElement("tr");
        row.innerHTML = `
          <td>${expense.amount}</td>
          <td>${expense.description}</td>
          <td>${expense.category}</td>
          <td><button data-index="${index}" class="delete-btn">Delete</button></td>
        `;
        expensesTableBody.appendChild(row);
      });
    };
  
    const renderSummary = () => {
        const totals = expenses.reduce((acc, curr) => {
          acc[curr.category] = (acc[curr.category] || 0) + Number(curr.amount);
          return acc;
        }, {});
      
        categorySummary.innerHTML = "";
        for (const [category, total] of Object.entries(totals)) {
          const li = document.createElement("li");
          li.classList.add("summary-item");
          li.innerHTML = `<span class="category">${category}</span><span class="total">₹${total}</span>`;
          categorySummary.appendChild(li);
        }
      };
      
  
    const updateChart = () => {
      const ctx = chartContainer.getContext("2d");
      const totals = expenses.reduce((acc, curr) => {
        acc[curr.category] = (acc[curr.category] || 0) + Number(curr.amount);
        return acc;
      }, {});
      new Chart(ctx, {
        type: "pie",
        data: {
          labels: Object.keys(totals),
          datasets: [{
            data: Object.values(totals),
            backgroundColor: ["#FF6384", "#36A2EB", "#FFCE56", "#4CAF50", "#FF2E7E", "#D52DB7"]
          }]
        }
      });
    };
  
    expenseForm.addEventListener("submit", (e) => {
      e.preventDefault();
      const amount = document.getElementById("amount").value;
      const description = document.getElementById("description").value;
      const category = document.getElementById("category").value;
  
      expenses.push({ amount, description, category });
      localStorage.setItem("expenses", JSON.stringify(expenses));
      renderExpenses();
      renderSummary();
      updateChart();
  
      expenseForm.reset();
    });
  
    expensesTableBody.addEventListener("click", (e) => {
      if (e.target.classList.contains("delete-btn")) {
        const index = e.target.getAttribute("data-index");
        expenses.splice(index, 1);
        localStorage.setItem("expenses", JSON.stringify(expenses));
        renderExpenses();
        renderSummary();
        updateChart();
      }
    });
  
    renderExpenses();
    renderSummary();
    updateChart();
  });
  