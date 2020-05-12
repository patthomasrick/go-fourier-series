set title "Fourier Approximation"
plot "plot_sin.dat" using 1:2 with lines t "Fourier Approximation", '' using 1:3 with lines t "Original Function"
pause -1