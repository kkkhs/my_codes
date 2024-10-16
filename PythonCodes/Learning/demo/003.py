import matplotlib.pyplot as plt
import numpy as np
t = np.linspace(0,2*np.pi,1000)
x = 16*np.sin(t)**3
y = 13*np.cos(t) - 5*np.cos(2*t) – 2*np.cos(3*t) – np.cos(4*t)
fig = plt.figure()
ax1 = fig.add_subplot(2,3,1)
ax2 = fig.add_subplot(2,3,2)
ax3 = fig.add_subplot(2,3,3)
ax4 = fig.add_subplot(2,3,4)
ax5 = fig.add_subplot(2,3,5)
ax6 = fig.add_subplot(2,3,6)
ax1.plot(x,y)
ax2.plot([2,0,2,2,2,1,8,0,2,7])
ax3.plot(x,y)
ax4.plot(x,y)
ax5.plot(x,y)
ax6.plot(x,y)
